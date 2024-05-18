package dto

type RegisterPatient struct {
	IdentityNumberInt   int64  `json:"identityNumber"`
	IdentityNumberStr   string `validate:"required,min=16,max=16"`
	PhoneNumber         string `json:"phoneNumber" validate:"required,startswith=+62,min=10,max=15"`
	Name                string `json:"name" validate:"required,min=5,max=30"`
	BirthDate           string `json:"birthDate" validate:"required,datetime=2006-01-02T15:04:05.000Z"`
	Gender              string `json:"gender" validate:"required,oneof=male female"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,ValidateURL"`
}

type PatientFilter struct {
	IdentityNumber int64  `json:"identityNumber"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Name           string `json:"name"`
	PhoneNumber    int64  `json:"phoneNumber"`
	CreatedAt      string `json:"createdAt"`
}

type CreateRecord struct {
	IdentityNumberInt int64  `json:"identityNumber" `
	IdentityNumberStr string `validate:"required,min=16,max=16"`
	Symptoms          string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications       string `json:"medications" validate:"required,min=1,max=2000"`
	CreatedBy         string
}

type RecordFilter struct {
	IdentityNumberStr string
	CreatedBy         struct {
		UserID string `json:"createdBy.userId"`
		NIP    string `json:"createdBy.nip"`
	} `json:"createdBy"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	CreatedAt string `json:"createdAt"`
}
