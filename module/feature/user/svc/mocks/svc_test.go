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

func TestUserSvc_CheckUserByIdAndRole_Error(t *testing.T) {
	mockCtrl, mockUserRepo, userService := SetUpDependencies(t)
	defer mockCtrl.Finish()

	id := "test_id"
	role := "test_role"

	Convey("When id and role is not found", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(id, role).Return(false, nil)
		exists, err := userService.CheckUserByIdAndRole(id, role)
		So(err, ShouldBeNil)
		So(exists, ShouldBeFalse)
	})

	Convey("When id and role is found", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(id, role).Return(true, nil)
		exists, err := userService.CheckUserByIdAndRole(id, role)
		So(err, ShouldBeNil)
		So(exists, ShouldBeTrue)
	})

	Convey("When internal server error", t, func() {
		mockUserRepo.EXPECT().CheckUserByIdAndRole(id, role).Return(false, response.NewInternalServerError(""))
		exists, err := userService.CheckUserByIdAndRole(id, role)
		So(err, ShouldNotBeNil)
		So(exists, ShouldBeFalse)
		So(errors.Is(err, response.NewInternalServerError("")), ShouldNotBeNil)
	})
}
