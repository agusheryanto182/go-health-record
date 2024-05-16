package entities

import (
	"database/sql"
)

type User struct {
	Id                  string         `db:"id"`
	Nip                 int64          `db:"nip"`
	Name                string         `db:"name"`
	Password            sql.NullString `db:"password"`
	Role                string         `db:"role"`
	IdentityCardScanImg sql.NullString `db:"identity_card_scan_img"`
	DeletedAt           sql.NullString `db:"deleted_at"`
	CreatedAt           string         `db:"created_at"`
}

var Role = struct {
	IT    string
	Nurse string
}{
	IT:    "it",
	Nurse: "nurse",
}
