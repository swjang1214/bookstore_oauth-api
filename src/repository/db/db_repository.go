package db

import (
	"fmt"

	"github.com/swjang1214/bookstore_oauth-api/src/db/mysql_db"
	"github.com/swjang1214/bookstore_oauth-api/src/domain/access_token"
	"github.com/swjang1214/bookstore_oauth-api/src/logger"
	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

const (
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type IDBRepository interface {
	Create(*access_token.AccessToken) *errors.RestError
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	UpdateExpirationTime(*access_token.AccessToken) *errors.RestError
}
type dbRepository struct {
}

func NewRepository() IDBRepository {
	return &dbRepository{}
}

func (*dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {

	stmt, err := mysql_db.Client.Prepare(queryGetAccessToken)
	if err != nil {
		logger.Error("error when trying to prepare queryGetAccessToken statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	if result == nil {
		logger.Error(fmt.Sprintf("error when trying to QueryRow by %s", id), nil)
		return nil, errors.NewInternalServerError("database error")
	}
	var at access_token.AccessToken

	if err := result.Scan(&at.AccessToken, &at.UserId, &at.ClientId, &at.Expires); err != nil {
		logger.Error("error when trying to get AccessToken by id", err)
		return nil, errors.NewInternalServerError("database error")
	}

	return &at, nil
}
func (*dbRepository) Create(token *access_token.AccessToken) *errors.RestError {

	stmt, err := mysql_db.Client.Prepare(queryCreateAccessToken)
	if err != nil {
		logger.Error("error when trying to prepare queryGetAccessToken statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(token.AccessToken, token.UserId, token.ClientId, token.Expires)
	if err != nil {
		logger.Error("error when trying to save token", err)
		return errors.NewInternalServerError("database error")
	}

	_, err = insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new token", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (*dbRepository) UpdateExpirationTime(token *access_token.AccessToken) *errors.RestError {
	stmt, err := mysql_db.Client.Prepare(queryUpdateExpires)
	if err != nil {
		logger.Error("error when trying to prepare queryUpdateExpires statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(token.Expires, token.AccessToken)
	if err != nil {
		logger.Error("error when trying to update token", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
