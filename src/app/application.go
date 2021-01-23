package app

import (
	"github.com/gin-gonic/gin"
	"github.com/swjang1214/bookstore_oauth-api/src/domain/access_token"
	"github.com/swjang1214/bookstore_oauth-api/src/handlers"
	"github.com/swjang1214/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := handlers.NewHttpHandler(
		access_token.NewSevice(rest.NewRestUsersRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	//router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)
	router.Run(":8000")

}
