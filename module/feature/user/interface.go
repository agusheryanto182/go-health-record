package user

import (
	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	"github.com/gofiber/fiber/v2"
)

type UserRepoInterface interface {
	Register(payload *entities.User) (string, error)
	IsNipExist(nip int64) (bool, error)
	GetUser(nip int64, role string) (*entities.User, error)
	GetUserByFilters(filters *dto.UserFilter) ([]*dto.UserFilterResponses, error)
	GetUserByID(id string) (*entities.User, error)
	UpdateUserNurse(req *dto.UpdateUserNurse) error
	DeleteUserNurse(req *dto.DeleteUserNurse) error
	SetPasswordNurse(req *dto.SetPasswordNurse) error
	CheckUserByIdAndRole(id string, role string) (bool, error)
}

type UserSvcInterface interface {
	Register(req *dto.RegisterUser) (*dto.RegisterAndLoginUserResponse, error)
	Login(req *dto.LoginUser) (*dto.RegisterAndLoginUserResponse, error)
	GetUserByFilters(filters *dto.UserFilter) ([]*dto.UserFilterResponses, error)
	GetUserByID(id string) (*entities.User, error)
	UpdateUserNurse(req *dto.UpdateUserNurse) error
	DeleteUserNurse(req *dto.DeleteUserNurse) error
	SetPasswordNurse(req *dto.SetPasswordNurse) error
	CheckUserByIdAndRole(id string, role string) (bool, error)
}

type UserHandlerInterface interface {
	RegisterIt(c *fiber.Ctx) error
	LoginIt(c *fiber.Ctx) error
	LoginNurse(c *fiber.Ctx) error
	RegisterNurse(c *fiber.Ctx) error
	GetUserByFilters(c *fiber.Ctx) error
	UpdateUserNurse(c *fiber.Ctx) error
	DeleteUserNurse(c *fiber.Ctx) error
	SetPasswordNurse(c *fiber.Ctx) error
}
