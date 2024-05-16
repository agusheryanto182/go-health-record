package hash

import (
	"github.com/agusheryanto182/go-health-record/config"
	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type Hash struct {
	cfg *config.Global
}

func NewHash(cfg *config.Global) HashInterface {
	return &Hash{
		cfg: cfg,
	}
}

func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cfg.Bcrypt.Salt)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
