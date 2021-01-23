package services

import (
	"strings"

	"github.com/swjang1214/bookstore_oauth-api/src/models/access_token"
	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

// type IRepository interface {
// 	GetById(string) (*AccessToken, *errors.RestError)
// 	Create(*AccessToken) *errors.RestError
// 	UpdateExpirationTime(*AccessToken) *errors.RestError
// }

type IService interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(*AccessToken) *errors.RestError
	UpdateExpirationTime(*AccessToken) *errors.RestError
}

type service struct {
	repository IRepository
}

func NewService(repo IRepository) IService {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
func (s *service) Create(token *AccessToken) *errors.RestError {
	// token.AccessToken = strings.TrimSpace(token.AccessToken)
	// if len(token.AccessToken) == 0 {
	// 	return errors.NewBadRequestError("invalid access token id")
	// }

	err := token.Validate()
	if err != nil {
		return err
	}

	return s.repository.Create(token)
}
func (s *service) UpdateExpirationTime(token *AccessToken) *errors.RestError {

	err := token.Validate()
	if err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(token)
}
