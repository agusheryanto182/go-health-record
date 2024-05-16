package svc

import (
	"database/sql"

	"github.com/agusheryanto182/go-health-record/module/entities"
	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/module/feature/user/dto"
	"github.com/agusheryanto182/go-health-record/utils/hash"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/go-playground/validator/v10"
)

type UserSvc struct {
	userRepo  user.UserRepoInterface
	validator *validator.Validate
	hash      hash.HashInterface
	jwtSvc    jwt.JWTInterface
}

// CheckUserByIdAndRole implements user.UserSvcInterface.
func (u *UserSvc) CheckUserByIdAndRole(id string, role string) (bool, error) {
	exists, err := u.userRepo.CheckUserByIdAndRole(id, role)
	if err != nil {
		return false, response.NewInternalServerError("errors when check user by id and role" + err.Error())
	}
	return exists, nil
}

// DeleteUserNurse implements user.UserSvcInterface.
func (u *UserSvc) DeleteUserNurse(req *dto.DeleteUserNurse) error {
	if err := u.userRepo.DeleteUserNurse(req); err != nil {
		return err
	}
	return nil
}

// SetPasswordNurse implements user.UserSvcInterface.
func (u *UserSvc) SetPasswordNurse(req *dto.SetPasswordNurse) error {
	_, err := u.userRepo.CheckUserByIdAndRole(req.ID, req.Role)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return response.NewBadRequestError("user id is not nurse")
		}
		return response.NewInternalServerError("errors when check user by id and role" + err.Error())
	}

	if err := u.validator.Struct(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	hashed, err := u.hash.HashPassword(req.Password)
	if err != nil {
		return response.NewInternalServerError("errors when generate hashed password" + err.Error())
	}

	req.Password = hashed
	if err := u.userRepo.SetPasswordNurse(req); err != nil {
		return err
	}
	return nil
}

// UpdateUserNurse implements user.UserSvcInterface.
func (u *UserSvc) UpdateUserNurse(req *dto.UpdateUserNurse) error {
	if err := u.validator.Struct(req); err != nil {
		return response.NewBadRequestError(err.Error())
	}

	if err := u.userRepo.UpdateUserNurse(req); err != nil {
		return err
	}
	return nil
}

// GetUserByID implements user.UserSvcInterface.
func (u *UserSvc) GetUserByID(id string) (*entities.User, error) {
	user := new(entities.User)
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, response.NewNotFoundError("user not found")
		}
		return nil, response.NewInternalServerError("errors when get user by id" + err.Error())
	}

	return user, nil
}

// GetUserByFilters implements user.UserSvcInterface.
func (u *UserSvc) GetUserByFilters(filters *dto.UserFilter) ([]*dto.UserFilterResponses, error) {
	result, err := u.userRepo.GetUserByFilters(filters)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return []*dto.UserFilterResponses{}, nil
		}
		return nil, response.NewInternalServerError("errors when get user by filters" + err.Error())
	}

	return result, nil
}

// Login implements user.UserSvcInterface.
func (u *UserSvc) Login(req *dto.LoginUser) (*dto.RegisterAndLoginUserResponse, error) {
	if req.Role == entities.Role.IT {
		if err := u.validator.Var(req.Nip, "required,ValidateNipIt"); err != nil {
			return nil, response.NewBadRequestError("NIP is not valid")
		}
	} else {
		if err := u.validator.Var(req.Nip, "required,ValidateNipNurse"); err != nil {
			return nil, response.NewBadRequestError("NIP is not valid")
		}
	}

	if err := u.validator.Struct(req); err != nil {
		return nil, response.NewBadRequestError(err.Error())
	}

	user, err := u.userRepo.GetUser(req.Nip, req.Role)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, response.NewNotFoundError("user not found")
		}
		return nil, response.NewInternalServerError("errors when get user" + err.Error())
	}

	if user.Role != req.Role {
		return nil, response.NewNotFoundError("user is not found")
	}

	if user.Password.String == "" {
		return nil, response.NewBadRequestError("user is not having access")
	}

	if !u.hash.CheckPasswordHash(req.Password, user.Password.String) {
		return nil, response.NewBadRequestError("password not match")
	}

	token, err := u.jwtSvc.GenerateJWT(user.Id, user.Role)
	if err != nil {
		return nil, response.NewInternalServerError("errors when generate token" + err.Error())
	}

	return &dto.RegisterAndLoginUserResponse{
		ID:          user.Id,
		Nip:         user.Nip,
		Name:        user.Name,
		AccessToken: token,
	}, nil
}

// Register implements user.UserSvcInterface.
func (u *UserSvc) Register(req *dto.RegisterUser) (*dto.RegisterAndLoginUserResponse, error) {
	if req.Role == entities.Role.IT {
		if err := u.validator.Var(req.Nip, "required,ValidateNipIt"); err != nil {
			return nil, response.NewBadRequestError("NIP is not valid : " + err.Error())
		}
		if err := u.validator.Var(req.Password, "required,min=5,max=33"); err != nil {
			return nil, response.NewBadRequestError("Password is not valid")
		}
	} else {
		if err := u.validator.Var(req.Nip, "required,ValidateNipNurse"); err != nil {
			return nil, response.NewBadRequestError("NIP is not valid : " + err.Error())
		}
		if err := u.validator.Var(req.IdentityCardScanImg, "required,ValidateURL"); err != nil {
			return nil, response.NewBadRequestError("Identity Card Scan Image is not valid")
		}
	}

	if err := u.validator.Struct(req); err != nil {
		return nil, response.NewBadRequestError(err.Error())
	}

	isNipExists, err := u.userRepo.IsNipExist(req.Nip)
	if err != nil {
		return nil, response.NewInternalServerError("errors when check nip" + err.Error())
	}

	if isNipExists {
		return nil, response.NewConflictError("NIP already exists")
	}

	var hashed string
	var newPassword sql.NullString
	var newIdentityCardScanImg sql.NullString
	if req.Role == entities.Role.IT {
		hashed, err = u.hash.HashPassword(req.Password)
		if err != nil {
			return nil, response.NewInternalServerError("errors when generate hashed password" + err.Error())
		}
		newPassword = sql.NullString{String: hashed, Valid: true}
		newIdentityCardScanImg = sql.NullString{String: "", Valid: false}
	} else {
		newIdentityCardScanImg = sql.NullString{String: req.IdentityCardScanImg, Valid: true}
		newPassword = sql.NullString{String: "", Valid: false}
	}

	payload := &entities.User{
		Nip:                 req.Nip,
		Name:                req.Name,
		Password:            newPassword,
		Role:                req.Role,
		IdentityCardScanImg: newIdentityCardScanImg,
	}

	id, err := u.userRepo.Register(payload)
	if err != nil {
		return nil, response.NewInternalServerError("errors when register user" + err.Error())
	}

	var token string
	if req.Role == entities.Role.IT {
		token, err = u.jwtSvc.GenerateJWT(id, payload.Name)
		if err != nil {
			return nil, response.NewInternalServerError("errors when generate token" + err.Error())
		}
	}
	return &dto.RegisterAndLoginUserResponse{
		ID:          id,
		Nip:         req.Nip,
		Name:        req.Name,
		AccessToken: token,
	}, nil
}

func NewUserService(userRepo user.UserRepoInterface, validator *validator.Validate, hash hash.HashInterface, jwtSvc jwt.JWTInterface) user.UserSvcInterface {
	return &UserSvc{
		userRepo:  userRepo,
		validator: validator,
		hash:      hash,
		jwtSvc:    jwtSvc,
	}
}
