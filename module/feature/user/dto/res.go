package dto

type RegisterAndLoginUserResponse struct {
	ID          string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken,omitempty"`
}

type UserFilterResponses struct {
	Id        string `json:"userId" db:"id"`
	Nip       int64  `json:"nip" db:"nip"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}
