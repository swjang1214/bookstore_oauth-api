package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	fmt.Println("start test cases ...")
	//rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"swjang1214@gmail.com", "password":"ksj0208##"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := restUserRepository{}
	user, err := repository.LoginUser("swjang1214@gmail.com", "ksj0208##")
	fmt.Println(user)
	fmt.Println(err)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}
func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
