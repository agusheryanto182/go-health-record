package svc

import (
	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/medical"
	"github.com/agusheryanto182/go-health-record/module/feature/medical/dto"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/go-playground/validator/v10"
)

type MedicalSvc struct {
	repo      medical.MedicalRepoInterface
	validator *validator.Validate
}

// CreateRecord implements medical.MedicalSvcInterface.
func (m *MedicalSvc) CreateRecord(req *dto.CreateRecord) error {
	if err := m.validator.Struct(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	if err := m.repo.CreateRecord(req); err != nil {
		return err
	}

	return nil
}

// GetPatientByFilters implements medical.MedicalSvcInterface.
func (m *MedicalSvc) GetPatientByFilters(filters *dto.PatientFilter) ([]*entities.Patient, error) {
	patients, err := m.repo.GetPatientByFilters(filters)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return []*entities.Patient{}, nil
		}
		return nil, response.NewInternalServerError("errors when get patient by filters : " + err.Error())
	}

	return patients, nil
}

// GetRecordByFilters implements medical.MedicalSvcInterface.
func (m *MedicalSvc) GetRecordByFilters(filters *dto.RecordFilter) ([]*dto.RecordResponses, error) {
	records, err := m.repo.GetRecordByFilters(filters)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return []*dto.RecordResponses{}, nil
		}
		return nil, response.NewInternalServerError("errors when get record by filters : " + err.Error())
	}

	return records, nil
}

// RegisterPatient implements medical.MedicalSvcInterface.
func (m *MedicalSvc) RegisterPatient(req *dto.RegisterPatient) error {
	if err := m.validator.Struct(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	if err := m.repo.RegisterPatient(req); err != nil {
		return err
	}

	return nil
}

func NewMedicalSvc(repo medical.MedicalRepoInterface, validator *validator.Validate) medical.MedicalSvcInterface {
	return &MedicalSvc{repo: repo, validator: validator}
}
