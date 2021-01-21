package db

import (
	"github.com/swjang1214/bookstore_oauth-api/src/domain/access_token"
	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

type IDBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}
type dbRepository struct {
}

func NewRepository() IDBRepository {
	return &dbRepository{}
}

func (*dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database error")
}
