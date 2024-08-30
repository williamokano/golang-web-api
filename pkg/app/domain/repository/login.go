package repository

import "github.com/williamokano/golang-web-api/pkg/app/domain/entities"

type LoginRepository interface {
	FindByUsername(username string) (*entities.Login, error)
}
