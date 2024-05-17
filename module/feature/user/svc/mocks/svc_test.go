package svc_test

import (
	"errors"
	"testing"

	"github.com/agusheryanto182/go-health-record/config"
	repo "github.com/agusheryanto182/go-health-record/module/feature/user/repo/mocks"
	"github.com/agusheryanto182/go-health-record/module/feature/user/svc"
	"github.com/agusheryanto182/go-health-record/utils/hash"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/agusheryanto182/go-health-record/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func SetUpCfg() *config.Global {
	cfg := &config.Global{}
	cfg.Database.DbName = "health-record"
	cfg.Database.Port = "5432"
	cfg.Database.Host = "localhost"
	cfg.Database.Username = "postgres"
	cfg.Database.Password = "postgres"
	cfg.Database.Params = "sslmode=disable"
	cfg.Jwt.Secret = "i-am-a-secret-key"
	cfg.Bcrypt.Salt = 8
	cfg.AWS.ID = "ID"
	cfg.AWS.SecretKey = "SECRET KEY"
	cfg.AWS.BucketName = "BUCKET NAME"
	cfg.AWS.Region = "REGION"

	return cfg
}

func SetUpDependencies(t *testing.T) (*gomock.Controller, *repo.MockUserRepoInterface, *svc.UserSvc) {
	ctrl := gomock.NewController(t)

	mockUserRepo := repo.NewMockUserRepoInterface(ctrl)

	mockValidator := validator.New()
	// register validation
	mockValidator.RegisterValidation("ValidateNipIt", validation.ValidateNipIt)
	mockValidator.RegisterValidation("ValidateNipNurse", validation.ValidateNipNurse)
	mockValidator.RegisterValidation("ValidateImage", validation.ValidateImage)
	mockValidator.RegisterValidation("ValidatePhoneNumber", validation.ValidatePhoneNumberFormat)
	mockValidator.RegisterValidation("ValidateURL", validation.ValidateURL)

	cfg := SetUpCfg()
	mockHash := hash.NewHash(cfg)
	mockJWT := jwt.NewJWTService(cfg)
	userService := svc.NewUserService(mockUserRepo, mockValidator, mockHash, mockJWT)

	return ctrl, mockUserRepo, userService.(*svc.UserSvc)

}

func TestUserSvc_SetPasswordNurse(t *testing.T) {
	mockCtrl, mockUserRepo, userService := SetUpDependencies(t)
	defer mockCtrl.Finish()

	Convey("When set password nurse success", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
		mockUserRepo.EXPECT().SetPasswordNurse(gomock.Any()).Return(nil)
		err := userService.SetPasswordNurse(nil)
		So(err, ShouldBeNil)
	})

	Convey("When user id is not nurse", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, nil)
		err := userService.SetPasswordNurse(nil)
		So(err, ShouldNotBeNil)
		So(errors.Is(err, response.NewBadRequestError("user id is not nurse")), ShouldBeTrue)
	})

	Convey("When internal server error", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, response.NewInternalServerError(""))
		err := userService.SetPasswordNurse(nil)
		So(err, ShouldNotBeNil)
		So(errors.Is(err, response.NewInternalServerError("")), ShouldBeTrue)
	})

	Convey("When validation error", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
		err := userService.SetPasswordNurse(nil)
		So(err, ShouldNotBeNil)
		So(errors.Is(err, response.NewBadRequestError("")), ShouldBeTrue)
	})

	Convey("When generate hashed password error", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
		mockUserRepo.EXPECT().SetPasswordNurse(gomock.Any()).Return(response.NewInternalServerError(""))
		err := userService.SetPasswordNurse(nil)
		So(err, ShouldNotBeNil)
		So(errors.Is(err, response.NewInternalServerError("")), ShouldBeTrue)
	})

}
func TestUserSvc_DeleteUserNurse(t *testing.T) {
	Convey("When delete user nurse", t, func() {
		mockCtrl, mockUserRepo, userService := SetUpDependencies(t)
		defer mockCtrl.Finish()

		Convey("Case delete user nurse success", func() {
			mockUserRepo.EXPECT().DeleteUserNurse(gomock.Any()).Return(nil)
			err := userService.DeleteUserNurse(nil)
			So(err, ShouldBeNil)
		})

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
	})
}

func TestUserSvc_CheckUserByIdAndRole_Error(t *testing.T) {
	Convey("When check user by id and role error", t, func() {
		mockCtrl, mockUserRepo, userService := SetUpDependencies(t)
		defer mockCtrl.Finish()

		Convey("Case user not found", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewNotFoundError(""))
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(errors.Is(err, response.NewNotFoundError("")), ShouldNotBeNil)
			So(exists, ShouldBeFalse)
		})

		Convey("Case user found", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(true, nil)
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(err, ShouldBeNil)
			So(exists, ShouldBeTrue)
		})

		Convey("Case internal server error", func() {
			mockUserRepo.EXPECT().CheckUserByIdAndRole(gomock.Any(), gomock.Any()).Return(false, response.NewInternalServerError(""))
			exists, err := userService.CheckUserByIdAndRole(gomock.Any().String(), gomock.Any().String())
			So(exists, ShouldBeFalse)
			So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
		})
	})

}
