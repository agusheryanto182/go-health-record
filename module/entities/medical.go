package entities

type Patient struct {
	IdentityNumber      int64  `json:"identityNumber" db:"identity_number"`
	PhoneNumber         string `json:"phoneNumber" db:"phone_number"`
	Name                string `json:"name" db:"name"`
	BirthDate           string `json:"birthDate" db:"birth_date"`
	Gender              string `json:"gender" db:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty" db:"identity_card_scan_img"`
	CreatedAt           string `json:"createdAt" db:"created_at"`
}

type Record struct {
	ID             string `json:"id" db:"id"`
	IdentityNumber int64  `json:"identityNumber" db:"identity_number"`
	Symptoms       string `json:"symptoms" db:"symptoms"`
	Medications    string `json:"medications" db:"medications"`
	CreatedBy      string `json:"createdBy" db:"created_by"`
	CreatedAt      string `json:"createdAt" db:"created_at"`
}
