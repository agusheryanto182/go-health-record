package repo

import (
	"strconv"
	"strings"
	"time"

	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/medical"
	"github.com/agusheryanto182/go-health-record/module/feature/medical/dto"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/jmoiron/sqlx"
)

type MedicalRepo struct {
	db *sqlx.DB
}

// CreateRecord implements medical.MedicalRepoInterface.
func (m *MedicalRepo) CreateRecord(payload *dto.CreateRecord) error {
	query :=
		`
	WITH checkPatient AS (
		SELECT EXISTS (
			SELECT 1 FROM patients WHERE identity_number = $1
		) AS exists
	)
	INSERT INTO records (identity_number, symptoms, medications, created_by)
	SELECT $1, $2, $3, $4
	WHERE (SELECT exists FROM checkPatient) = true
	`

	res, err := m.db.Exec(query, payload.IdentityNumberInt, payload.Symptoms, payload.Medications, payload.CreatedBy)
	if err != nil {
		return response.NewInternalServerError("errors when create record : " + err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return response.NewBadRequestError("patient not found")
	}

	return nil
}

// GetPatientByFilters implements medical.MedicalRepoInterface.
func (m *MedicalRepo) GetPatientByFilters(filters *dto.PatientFilter) ([]*entities.Patient, error) {
	query :=
		`
	SELECT * FROM patients
	WHERE 1 = 1
	`

	params := make([]interface{}, 0)

	if filters.IdentityNumber > 0 {
		query += " AND identity_number = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.IdentityNumber)
	}

	nameCleaned := strings.ReplaceAll(filters.Name, "\"", "")
	if nameCleaned != "" {
		query += " AND name ILIKE '%' || $" + strconv.Itoa(len(params)+1) + " || '%'"
		params = append(params, nameCleaned)
	}

	if filters.PhoneNumber != 0 {
		query += " AND CAST(phone_number AS TEXT) LIKE CAST($" + strconv.Itoa(len(params)+1) + " AS TEXT) || '%'"
		params = append(params, "+"+strconv.Itoa(int(filters.PhoneNumber)))
	}

	createdAtCleaned := strings.ReplaceAll(filters.CreatedAt, "\"", "")
	if createdAtCleaned == "asc" || createdAtCleaned == "desc" {
		query += " ORDER BY created_at " + createdAtCleaned
	}

	if createdAtCleaned == "" {
		query += " ORDER BY created_at DESC"
	}

	if filters.Limit != 0 {
		query += " LIMIT $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Limit)
	} else {
		query += " LIMIT 5"
	}

	if filters.Offset != 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	} else {
		query += " OFFSET 0"
	}

	patients := []*entities.Patient{}

	rows, err := m.db.Queryx(query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		createdAtTemp := time.Time{}
		patient := new(entities.Patient)
		if err := rows.Scan(
			&patient.IdentityNumber,
			&patient.PhoneNumber,
			&patient.Name,
			&patient.BirthDate,
			&patient.Gender,
			&patient.IdentityCardScanImg,
			&createdAtTemp,
		); err != nil {
			return nil, err
		}

		patient.CreatedAt = createdAtTemp.Format(time.RFC3339)

		patients = append(patients, patient)
	}

	return patients, nil

}

// GetRecordByFilters implements medical.MedicalRepoInterface.
func (m *MedicalRepo) GetRecordByFilters(filters *dto.RecordFilter) ([]*dto.RecordResponses, error) {
	query :=
		`
	SELECT p.identity_number, p.phone_number, p.name, p.birth_date, p.gender, p.identity_card_scan_img,
	r.symptoms, r.medications, r.created_at,
	u.nip, u.name,u.id
	FROM records r
	JOIN patients p ON r.identity_number = p.identity_number
	JOIN users u ON r.created_by = u.id
	WHERE 1 = 1
	`

	params := make([]interface{}, 0)

	if filters.IdentityNumberStr != "" {
		query += " AND p.identity_number = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.IdentityNumberStr)
	}

	cleanedUserID := strings.ReplaceAll(filters.CreatedBy.UserID, "\"", "")
	if filters.CreatedBy.UserID != "" {
		query += " AND r.created_by = $" + strconv.Itoa(len(params)+1)
		params = append(params, cleanedUserID)
	}

	nipCleaned := strings.ReplaceAll(filters.CreatedBy.NIP, "\"", "")
	if filters.CreatedBy.NIP != "" {
		query += " AND u.nip = $" + strconv.Itoa(len(params)+1)
		params = append(params, nipCleaned)
	}

	createdAtCleaned := strings.ReplaceAll(filters.CreatedAt, "\"", "")
	if createdAtCleaned == "asc" || createdAtCleaned == "desc" {
		query += " ORDER BY r.created_at " + createdAtCleaned
	}

	if createdAtCleaned == "" {
		query += " ORDER BY r.created_at DESC"
	}

	if filters.Limit != 0 {
		query += " LIMIT $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Limit)
	} else {
		query += " LIMIT 5"
	}

	if filters.Offset != 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	} else {
		query += " OFFSET 0"
	}

	records := []*dto.RecordResponses{}

	rows, err := m.db.Queryx(query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		createdAtTemp := time.Time{}
		record := new(dto.RecordResponses)
		if err := rows.Scan(
			&record.IdentityDetail.IdentityNumber,
			&record.IdentityDetail.PhoneNumber,
			&record.IdentityDetail.Name,
			&record.IdentityDetail.BirthDate,
			&record.IdentityDetail.Gender,
			&record.IdentityDetail.IdentityCardScanImg,
			&record.Symptoms,
			&record.Medications,
			&createdAtTemp,
			&record.CreatedBy.NIP,
			&record.CreatedBy.Name,
			&record.CreatedBy.UserID,
		); err != nil {
			return nil, err
		}
		record.CreatedAt = createdAtTemp.Format(time.RFC3339)

		records = append(records, record)
	}
	return records, nil
}

// RegisterPatient implements medical.MedicalRepoInterface.
func (m *MedicalRepo) RegisterPatient(payload *dto.RegisterPatient) error {
	query :=
		`
		WITH checkIdentityNumber AS (
			SELECT EXISTS (
				SELECT 1 FROM patients WHERE identity_number = $1
			) AS exists
		) 
		INSERT INTO patients (identity_number, phone_number, name, birth_date, gender, identity_card_scan_img)
		SELECT $1, $2, $3, $4, $5, $6
		WHERE NOT EXISTS (SELECT 1 FROM checkIdentityNumber WHERE exists = true)		
	`

	res, err := m.db.Exec(query, payload.IdentityNumberInt, payload.PhoneNumber, payload.Name, payload.BirthDate, payload.Gender, payload.IdentityCardScanImg)
	if err != nil {
		return response.NewInternalServerError("errors when register patient : " + err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return response.NewConflictError("identity number already exist")
	}
	return nil
}

func NewMedicalRepo(db *sqlx.DB) medical.MedicalRepoInterface {
	return &MedicalRepo{db: db}
}
