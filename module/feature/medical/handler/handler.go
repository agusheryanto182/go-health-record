package handler

import (
	"strconv"

	"github.com/agusheryanto182/go-health-record/module/feature/medical"
	"github.com/agusheryanto182/go-health-record/module/feature/medical/dto"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

type MedicalHandler struct {
	svc medical.MedicalSvcInterface
}

// CreateRecord implements medical.MedicalHandlerInterface.
func (m *MedicalHandler) CreateRecord(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*jwt.JWTPayload)
	req := new(dto.CreateRecord)

	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.IdentityNumberStr = strconv.FormatInt(req.IdentityNumberInt, 10)

	req.CreatedBy = currentUser.Id

	if err := m.svc.CreateRecord(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

// GetPatientByFilters implements medical.MedicalHandlerInterface.
func (m *MedicalHandler) GetPatientByFilters(c *fiber.Ctx) error {
	filters := new(dto.PatientFilter)

	if err := c.QueryParser(filters); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	patients, err := m.svc.GetPatientByFilters(filters)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    patients,
	})
}

// GetRecordByFilters implements medical.MedicalHandlerInterface.
func (m *MedicalHandler) GetRecordByFilters(c *fiber.Ctx) error {
	filters := new(dto.RecordFilter)

	if err := c.QueryParser(filters); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	filters.IdentityNumberStr = c.Query("identityDetail.identityNumber")

	records, err := m.svc.GetRecordByFilters(filters)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    records,
	})
}

// RegisterPatient implements medical.MedicalHandlerInterface.
func (m *MedicalHandler) RegisterPatient(c *fiber.Ctx) error {
	req := new(dto.RegisterPatient)

	if err := c.BodyParser(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	req.IdentityNumberStr = strconv.FormatInt(req.IdentityNumberInt, 10)

	if err := m.svc.RegisterPatient(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func NewMedicalHandler(svc medical.MedicalSvcInterface) medical.MedicalHandlerInterface {
	return &MedicalHandler{svc: svc}
}
