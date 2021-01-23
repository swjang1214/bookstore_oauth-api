package access_token

import (
	"strings"
	"time"

	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func (at *AccessToken) IsExpired() bool {
	//! Before : 비교 시간이 인자로 주어진 시간보다 전이면 true, 아니면 false
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
