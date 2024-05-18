package repo

import (
	"strconv"
	"strings"
	"time"

	entity "github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// CheckUserByIdAndRole implements user.UserRepoInterface.
func (u *UserRepository) CheckUserByIdAndRole(id string, role string) (bool, error) {
	var exists bool
	if err := u.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2 AND deleted_at IS NULL)", id, role); err != nil {
		return false, err
	}
	return exists, nil
}

// DeleteUserNurse implements user.UserRepoInterface.
func (u *UserRepository) DeleteUserNurse(req *dto.DeleteUserNurse) error {
	res, err := u.db.Exec(`
        WITH checkUser AS (
            SELECT id
            FROM users
            WHERE id = $1 AND role = $2 AND deleted_at IS NULL
        )
        UPDATE users
        SET deleted_at = $3
        FROM checkUser
        WHERE users.id = checkUser.id
    `, req.ID, req.Role, time.Now())
	if err != nil {
		return response.NewInternalServerError("errors when delete user : " + err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return response.NewNotFoundError("user is not found or already deleted")
	}

	return nil
}

// SetPasswordNurse implements user.UserRepoInterface.
func (u *UserRepository) SetPasswordNurse(req *dto.SetPasswordNurse) error {
	res, err := u.db.Exec("UPDATE users SET password = $1 WHERE id = $2 AND role = $3 AND deleted_at IS NULL", req.Password, req.ID, req.Role)
	if err != nil {
		return response.NewInternalServerError("errors when set password nurse : " + err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return response.NewBadRequestError("user is not nurse")

	}
	return nil
}

func (u *UserRepository) UpdateUserNurse(req *dto.UpdateUserNurse) error {
	var userExists, nipExists bool
	err := u.db.QueryRow(`
        WITH checkData AS (
            SELECT
                EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2 AND deleted_at IS NULL) AS user_exists,
                EXISTS(SELECT 1 FROM users WHERE nip = $3 AND role = $2 AND id != $1 AND deleted_at IS NULL) AS nip_exists
        )
        SELECT user_exists, nip_exists FROM checkData
    `, req.ID, req.Role, req.Nip).Scan(&userExists, &nipExists)
	if err != nil {
		return err
	}

	if !userExists {
		return response.NewNotFoundError("user not found")
	}

	if nipExists {
		return response.NewNotFoundError("NIP already exists")
	}

	_, err = u.db.Exec(`
        UPDATE users
        SET nip = $1, name = $2
        WHERE id = $3 AND role = $4 AND deleted_at IS NULL
    `, req.Nip, req.Name, req.ID, req.Role)
	if err != nil {
		return err
	}

	return nil
}

// GetUser implements user.UserRepoInterface.
func (u *UserRepository) GetUser(nip int64, role string) (*entity.User, error) {
	user := new(entity.User)
	if err := u.db.Get(user, `
	SELECT * FROM users 
	WHERE nip = $1 
	AND role = $2 
	AND deleted_at IS NULL`,
		nip, role); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByFilters implements user.UserRepoInterface.
func (u *UserRepository) GetUserByFilters(filters *dto.UserFilter) ([]*dto.UserFilterResponses, error) {
	query := `
	SELECT id, nip, name, created_at 
	FROM users WHERE 1 = 1
	`

	params := make([]interface{}, 0)

	if filters.ID != "" {
		query += " AND id = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.ID)
	}

	cleanedName := strings.ReplaceAll(filters.Name, "\"", "")
	if cleanedName != "" {
		query += " AND name ILIKE '%' || $" + strconv.Itoa(len(params)+1) + " || '%'"
		params = append(params, cleanedName)
	}

	if filters.Nip != 0 {
		query += " AND CAST(nip AS TEXT) LIKE CAST($" + strconv.Itoa(len(params)+1) + " AS TEXT) || '%'"
		params = append(params, strconv.Itoa(int(filters.Nip)))
	}

	roleCleaned := strings.ReplaceAll(filters.Role, "\"", "")
	if roleCleaned != "" {
		if roleCleaned == "it" || roleCleaned == "nurse" {
			query += " AND role = $" + strconv.Itoa(len(params)+1)
			params = append(params, roleCleaned)
		}
	}

	query += " AND deleted_at IS NULL"

	createdAtCleaned := strings.ReplaceAll(filters.CreatedAt, "\"", "")
	if createdAtCleaned != "" {
		if createdAtCleaned == "asc" {
			query += " ORDER BY created_at ASC"
		} else if createdAtCleaned == "desc" {
			query += " ORDER BY created_at DESC"
		}
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

	rows, err := u.db.Queryx(query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*dto.UserFilterResponses{}
	for rows.Next() {
		var tempCreatedAt time.Time
		user := new(dto.UserFilterResponses)
		if err := rows.Scan(
			&user.Id,
			&user.Nip,
			&user.Name,
			&tempCreatedAt,
		); err != nil {
			return nil, err
		}

		user.CreatedAt = tempCreatedAt.Format(time.RFC3339)
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID implements user.UserRepoInterface.
func (u *UserRepository) GetUserByID(id string) (*entity.User, error) {
	user := new(entity.User)
	query := `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`
	if err := u.db.Get(user, query, id); err != nil {
		return nil, err
	}
	return user, nil
}

// IsNipExist implements user.UserRepoInterface.
func (u *UserRepository) IsNipExist(nip int64) (bool, error) {
	var exists bool
	if err := u.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE nip = $1 AND deleted_at IS NULL)", nip); err != nil {
		return false, err
	}
	return exists, nil
}

// Register implements user.UserRepoInterface.
func (u *UserRepository) Register(payload *entity.User) (string, error) {
	var id string

	query := `INSERT INTO users (nip, name, password, role, identity_card_scan_img) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := u.db.QueryRowx(query, payload.Nip, payload.Name, payload.Password, payload.Role, payload.IdentityCardScanImg.String)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func NewUserRepository(db *sqlx.DB) user.UserRepoInterface {
	return &UserRepository{db: db}
}
