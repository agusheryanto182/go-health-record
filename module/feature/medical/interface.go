package medical

import (
	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/medical/dto"
	"github.com/gofiber/fiber/v2"
)

type MedicalRepoInterface interface {
	RegisterPatient(payload *dto.RegisterPatient) error
	GetPatientByFilters(filters *dto.PatientFilter) ([]*entities.Patient, error)
	CreateRecord(payload *dto.CreateRecord) error
	GetRecordByFilters(filters *dto.RecordFilter) ([]*dto.RecordResponses, error)
}

type MedicalSvcInterface interface {
	RegisterPatient(req *dto.RegisterPatient) error
	GetPatientByFilters(filters *dto.PatientFilter) ([]*entities.Patient, error)
	CreateRecord(req *dto.CreateRecord) error
	GetRecordByFilters(filters *dto.RecordFilter) ([]*dto.RecordResponses, error)
}

type MedicalHandlerInterface interface {
	RegisterPatient(c *fiber.Ctx) error
	GetPatientByFilters(c *fiber.Ctx) error
	CreateRecord(c *fiber.Ctx) error
	GetRecordByFilters(c *fiber.Ctx) error
}
