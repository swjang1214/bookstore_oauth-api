package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/swjang1214/bookstore_oauth-api/src/domain/users"
	"github.com/swjang1214/bookstore_oauth-api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type IRestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}
type restUserRepository struct {
}

func NewRepository() IRestUserRepository {
	return &restUserRepository{}
}
func (*restUserRepository) LoginUser(email string, pwd string) (*users.User, *errors.RestError) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: pwd,
	}

	response := usersRestClient.Post("/users/login", req)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
