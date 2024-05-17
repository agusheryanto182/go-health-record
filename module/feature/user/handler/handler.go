package handler

import (
	"fmt"
	"strings"

	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSvc user.UserSvcInterface
}

func getCurrentUserIT(c *fiber.Ctx) (*jwt.JWTPayload, error) {
	currentUser := c.Locals("CurrentUser").(*jwt.JWTPayload)
	if currentUser.Role != entities.Role.IT {
		return nil, response.NewUnauthorizedError("Access denied: user not allowed to access this feature")
	}
	return currentUser, nil
}

// DeleteUserNurse implements user.UserHandlerInterface.
func (u *UserHandler) DeleteUserNurse(c *fiber.Ctx) error {
	if _, err := getCurrentUserIT(c); err != nil {
		return err
	}

	req := new(dto.DeleteUserNurse)

	req.ID = c.Params("userId", req.ID)
	req.Role = entities.Role.Nurse

	if err := u.userSvc.DeleteUserNurse(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

// SetPasswordNurse implements user.UserHandlerInterface.
func (u *UserHandler) SetPasswordNurse(c *fiber.Ctx) error {
	if _, err := getCurrentUserIT(c); err != nil {
		return err
	}

	req := new(dto.SetPasswordNurse)

	req.ID = c.Params("userId", req.ID)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}
	req.Role = entities.Role.Nurse

	if err := u.userSvc.SetPasswordNurse(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

// UpdateUserNurse implements user.UserHandlerInterface.
func (u *UserHandler) UpdateUserNurse(c *fiber.Ctx) error {
	if _, err := getCurrentUserIT(c); err != nil {
		return err
	}

	req := new(dto.UpdateUserNurse)

	req.ID = c.Params("userId", req.ID)
	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}
	req.Role = entities.Role.Nurse

	if err := u.userSvc.UpdateUserNurse(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

// GetUserByFilters implements user.UserHandlerInterface.
func (u *UserHandler) GetUserByFilters(c *fiber.Ctx) error {
	if _, err := getCurrentUserIT(c); err != nil {
		return err
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

	fmt.Println(req)
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
	if _, err := getCurrentUserIT(c); err != nil {
		return err
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
