package dto

type UserFilter struct {
	ID        string
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Name      string `json:"name"`
	Nip       int64  `json:"nip"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

type LoginUser struct {
	Nip      int64  `json:"nip"`
	Password string `json:"password" validate:"required,min=5,max=33"`
	Role     string
}

type RegisterUser struct {
	Nip                 int64  `json:"nip"`
	Name                string `json:"name" validate:"required,min=5,max=50"`
	Password            string `json:"password"`
	Role                string
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

type UpdateUserNurse struct {
	ID   string
	Nip  int64  `json:"nip" validate:"required,ValidateNipNurse"`
	Name string `json:"name" validate:"required,min=5,max=50"`
	Role string
}

type SetPasswordNurse struct {
	ID       string
	Password string `json:"password" validate:"required,min=5,max=33"`
	Role     string
}

type DeleteUserNurse struct {
	ID   string
	Role string
}
