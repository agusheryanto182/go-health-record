package svc_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/agusheryanto182/go-health-record/module/entities"
	dto "github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	repo "github.com/agusheryanto182/go-health-record/module/feature/user/repo/mocks"
	"github.com/agusheryanto182/go-health-record/module/feature/user/svc"
	hash "github.com/agusheryanto182/go-health-record/utils/hash/mocks"
	jwt "github.com/agusheryanto182/go-health-record/utils/jwt/mocks"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/agusheryanto182/go-health-record/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func SetUpDependencies(t *testing.T) (*gomock.Controller, *repo.MockUserRepoInterface, *svc.UserSvc, *hash.MockHashInterface, *jwt.MockJWTInterface) {
	ctrl := gomock.NewController(t)

	mockUserRepo := repo.NewMockUserRepoInterface(ctrl)

	mockValidator := validator.New()
	// register validation
	mockValidator.RegisterValidation("ValidateNipIt", validation.ValidateNipIt)
	mockValidator.RegisterValidation("ValidateNipNurse", validation.ValidateNipNurse)
	mockValidator.RegisterValidation("ValidateImage", validation.ValidateImage)
	mockValidator.RegisterValidation("ValidatePhoneNumber", validation.ValidatePhoneNumberFormat)
	mockValidator.RegisterValidation("ValidateURL", validation.ValidateURL)

	mockHash := hash.NewMockHashInterface(ctrl)
	mockJWT := jwt.NewMockJWTInterface(ctrl)
	userService := svc.NewUserService(mockUserRepo, mockValidator, mockHash, mockJWT)

	return ctrl, mockUserRepo, userService.(*svc.UserSvc), mockHash, mockJWT

}

func TestUserSvc_Login(t *testing.T) {
	Convey("When login", t, func() {
		mockCtrl, mockUserRepo, userService, mockHash, mockJWT := SetUpDependencies(t)
		defer mockCtrl.Finish()

		positifCase := &dto.LoginUser{
			Nip:      615220010298712,
			Password: "password123",
			Role:     "it",
		}

		resGetUser := &entities.User{
			Id:       "123456789",
			Nip:      615220010298712,
			Name:     "Suga",
			Role:     "it",
			Password: sql.NullString{String: "passwordHashed", Valid: true},
		}

		resGetUserNurseNotHavingAccess := &entities.User{
			Id:       "123456789",
			Nip:      615220010298712,
			Name:     "Suga",
			Role:     "it",
			Password: sql.NullString{String: "", Valid: false},
		}

		negatifCase := &dto.LoginUser{}

		Convey("Case request does not pass validation", func() {
			res, err := userService.Login(negatifCase)
			So(errors.Is(err, response.NewBadRequestError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case internal server error, when get user", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, response.NewInternalServerError(""))
			res, err := userService.Login(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case user not found", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(nil, response.NewNotFoundError(""))
			res, err := userService.Login(positifCase)
			So(errors.Is(err, response.NewNotFoundError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case password not match", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(resGetUser, nil)
			mockHash.EXPECT().CheckPasswordHash(gomock.Any(), gomock.Any()).Return(false)
			res, err := userService.Login(positifCase)
			So(errors.Is(err, response.NewBadRequestError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
			So(resGetUser.Nip, ShouldEqual, positifCase.Nip)
			So(resGetUser.Role, ShouldEqual, positifCase.Role)
			So(resGetUser.Password.String, ShouldNotBeNil)
		})

		Convey("Case user is not having access", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(resGetUserNurseNotHavingAccess, nil)
			res, err := userService.Login(positifCase)
			So(errors.Is(err, response.NewBadRequestError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
			So(resGetUserNurseNotHavingAccess.Nip, ShouldEqual, positifCase.Nip)
			So(resGetUserNurseNotHavingAccess.Role, ShouldEqual, positifCase.Role)
			So(resGetUserNurseNotHavingAccess.Password.String, ShouldEqual, "")
		})

		Convey("Case internal server error, when generate jwt", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(resGetUser, nil)
			mockHash.EXPECT().CheckPasswordHash(gomock.Any(), gomock.Any()).Return(true)
			mockJWT.EXPECT().GenerateJWT(gomock.Any(), gomock.Any()).Return("", response.NewInternalServerError(""))
			res, err := userService.Login(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
			So(resGetUser.Nip, ShouldEqual, positifCase.Nip)
			So(resGetUser.Role, ShouldEqual, positifCase.Role)
			So(resGetUser.Password.String, ShouldNotBeNil)
		})

		Convey("Case success", func() {
			mockUserRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(resGetUser, nil)
			mockHash.EXPECT().CheckPasswordHash(gomock.Any(), gomock.Any()).Return(true)
			mockJWT.EXPECT().GenerateJWT(gomock.Any(), gomock.Any()).Return("token", nil)
			res, err := userService.Login(positifCase)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.ID, ShouldEqual, resGetUser.Id)
			So(res.Nip, ShouldEqual, resGetUser.Nip)
			So(res.Name, ShouldEqual, resGetUser.Name)
			So(res.AccessToken, ShouldEqual, "token")
		})
	})
}

func TestUserSvc_RegisterUser(t *testing.T) {
	Convey("When register user", t, func() {
		mockCtrl, mockUserRepo, userService, mockHash, mockJWT := SetUpDependencies(t)
		defer mockCtrl.Finish()

		positifCase := &dto.RegisterUser{
			Nip:      615220010298712,
			Name:     "Suga 123",
			Role:     "it",
			Password: "password123",
		}

		negatifCase := &dto.RegisterUser{}

		Convey("Case request does not pass validation", func() {
			res, err := userService.Register(negatifCase)
			So(errors.Is(err, response.NewBadRequestError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case internal server error, when check is nip exist", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(false, response.NewInternalServerError(""))
			res, err := userService.Register(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case nip already exist", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(true, nil)
			res, err := userService.Register(positifCase)
			So(errors.Is(err, response.NewConflictError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case hashing password error", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(false, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("", response.NewInternalServerError(""))
			res, err := userService.Register(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case internal server error, when register user", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(false, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("hashed", nil)
			mockUserRepo.EXPECT().Register(gomock.Any()).Return("", response.NewInternalServerError(""))
			res, err := userService.Register(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case internal server error, when generate token", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(false, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("hashed", nil)
			mockUserRepo.EXPECT().Register(gomock.Any()).Return("id", nil)
			mockJWT.EXPECT().GenerateJWT(gomock.Any(), gomock.Any()).Return("", response.NewInternalServerError(""))
			res, err := userService.Register(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("Case success", func() {
			mockUserRepo.EXPECT().IsNipExist(gomock.Any()).Return(false, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("hashed", nil)
			mockUserRepo.EXPECT().Register(gomock.Any()).Return("id", nil)
			mockJWT.EXPECT().GenerateJWT(gomock.Any(), gomock.Any()).Return("token", nil)
			res, err := userService.Register(positifCase)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.ID, ShouldEqual, "id")
			So(res.Nip, ShouldEqual, positifCase.Nip)
			So(res.Name, ShouldEqual, positifCase.Name)
			So(res.AccessToken, ShouldEqual, "token")
		})
	})
}

func TestUserSvc_SetPasswordNurse(t *testing.T) {
	Convey("When set password nurse", t, func() {
		mockCtrl, mockUserRepo, userService, mockHash, _ := SetUpDependencies(t)
		defer mockCtrl.Finish()

		positifCase := &dto.SetPasswordNurse{
			ID:       "123456789",
			Role:     "it",
			Password: "password123",
		}

		negatifCase := &dto.SetPasswordNurse{
			Password: "123",
		}

		Convey("Case user not found", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewNotFoundError(""))
			err := userService.SetPasswordNurse(positifCase)
			So(errors.Is(err, response.NewNotFoundError("")), ShouldNotBeNil)
		})

		Convey("Case internal server error, when check user by id and role", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewInternalServerError(""))
			err := userService.SetPasswordNurse(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
		})

		Convey("Case request does not pass validation", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			err := userService.SetPasswordNurse(negatifCase)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, response.NewBadRequestError("")), ShouldNotBeNil)
		})

		Convey("Case hashing password error", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("", response.NewInternalServerError(""))
			err := userService.SetPasswordNurse(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
		})

		Convey("Case internal server error, when set password nurse", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("hashed", nil)
			mockUserRepo.EXPECT().SetPasswordNurse(gomock.Any()).Return(response.NewInternalServerError(""))
			err := userService.SetPasswordNurse(positifCase)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
		})

		Convey("Case success", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			mockHash.EXPECT().HashPassword(gomock.Any()).Return("hashed", nil)
			mockUserRepo.EXPECT().SetPasswordNurse(gomock.Any()).Return(nil)
			err := userService.SetPasswordNurse(positifCase)
			So(err, ShouldBeNil)
		})

	})

}
func TestUserSvc_DeleteUserNurse(t *testing.T) {
	Convey("When delete user nurse", t, func() {
		mockCtrl, mockUserRepo, userService, _, _ := SetUpDependencies(t)
		defer mockCtrl.Finish()

		Convey("Case internal server error", func() {
			mockUserRepo.EXPECT().DeleteUserNurse(gomock.Any()).Return(response.NewInternalServerError(""))
			err := userService.DeleteUserNurse(nil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldBeTrue)
		})

		Convey("Case user not found", func() {
			mockUserRepo.EXPECT().DeleteUserNurse(gomock.Any()).Return(response.NewNotFoundError(""))
			err := userService.DeleteUserNurse(nil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, response.NewNotFoundError("")), ShouldBeTrue)
		})

		Convey("Case delete user nurse success", func() {
			mockUserRepo.EXPECT().DeleteUserNurse(gomock.Any()).Return(nil)
			err := userService.DeleteUserNurse(nil)
			So(err, ShouldBeNil)
		})
	})
}

func TestUserSvc_CheckUserByIdAndRole_Error(t *testing.T) {
	Convey("When check user by id and role error", t, func() {
		mockCtrl, mockUserRepo, userService, _, _ := SetUpDependencies(t)
		defer mockCtrl.Finish()

		Convey("Case user not found", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewNotFoundError(""))
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(errors.Is(err, response.NewNotFoundError("")), ShouldNotBeNil)
			So(exists, ShouldBeFalse)
		})

		Convey("Case internal server error", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewInternalServerError(""))
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(exists, ShouldBeFalse)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
		})

		Convey("Case user found", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(err, ShouldBeNil)
			So(exists, ShouldBeTrue)
		})

	})

}
