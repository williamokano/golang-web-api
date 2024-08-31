package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/williamokano/golang-web-api/pkg/app/domain/entities"
	domainrepository "github.com/williamokano/golang-web-api/pkg/app/domain/repository"
)

type LoginRepository struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) *LoginRepository {
	return &LoginRepository{
		db: db,
	}
}

func (lp *LoginRepository) ListAll() ([]entities.Login, error) {
	return nil, nil
}

func (lp *LoginRepository) FindByUsername(username string) (*entities.Login, error) {
	var result entities.Login // Create an instance instead of a pointer
	err := lp.db.QueryRow("SELECT username, password, totp_code FROM login WHERE username = $1", username).Scan(
		&result.Username,
		&result.HashedPassword,
		&result.TOTPSecret,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("username %s not found", username)
		}
		return nil, err
	}
	return &result, nil // Return the address of the result
}

var _ domainrepository.LoginRepository = &LoginRepository{}
