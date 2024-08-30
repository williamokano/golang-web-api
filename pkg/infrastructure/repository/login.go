package repository

import (
	"database/sql"

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

func (lp *LoginRepository) FindByUsername(username string) (*entities.Login, error) {
	return nil, nil
}

var _ domainrepository.LoginRepository = &LoginRepository{}
