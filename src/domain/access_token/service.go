package access_token

import (
	"strings"

	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

type IRepository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type IService interface {
	GetById(string) (*AccessToken, *errors.RestError)
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
