package handler

import (
	"strings"

	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSvc user.UserSvcInterface
}

// GetUserByFilters implements user.UserHandlerInterface.
func (u *UserHandler) GetUserByFilters(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != entities.Role.IT {
		return response.NewUnauthorizedError("Access denied: user not allowed to access this feature")
	}

	req := new(dto.UserFilter)
	if err := c.QueryParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	userId := strings.ReplaceAll(c.Query("userId"), "\"", "")
	req.ID = userId

	result, err := u.userSvc.GetUserByFilters(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    result,
	})
}

// LoginIt implements user.UserHandlerInterface.
func (u *UserHandler) LoginIt(c *fiber.Ctx) error {
	req := new(dto.LoginUser)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.Role = entities.Role.IT
	resp, err := u.userSvc.Login(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

// LoginNurse implements user.UserHandlerInterface.
func (u *UserHandler) LoginNurse(c *fiber.Ctx) error {
	req := new(dto.LoginUser)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.Role = entities.Role.Nurse
	resp, err := u.userSvc.Login(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

// RegisterIt implements user.UserHandlerInterface.
func (u *UserHandler) RegisterIt(c *fiber.Ctx) error {
	req := new(dto.RegisterUser)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.Role = entities.Role.IT
	resp, err := u.userSvc.Register(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

// RegisterNurse implements user.UserHandlerInterface.
func (u *UserHandler) RegisterNurse(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entities.User)
	if currentUser.Role != entities.Role.IT {
		return response.NewUnauthorizedError("Access denied: user not allowed to access this feature")
	}

	req := new(dto.RegisterUser)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.Role = entities.Role.Nurse
	resp, err := u.userSvc.Register(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

func NewUserHandler(userSvc user.UserSvcInterface) user.UserHandlerInterface {
	return &UserHandler{userSvc: userSvc}
}
