package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swjang1214/bookstore_oauth-api/src/services"
)

type IHttpHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type httpHandler struct {
	service services.IService
}

func NewHttpHandler(service services.IService) IHttpHandler {
	return &httpHandler{
		service: service,
	}
}

func (handler *httpHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
func (handler *httpHandler) Create(c *gin.Context) {
	// var at access_token.AccessToken
	// if err := c.ShouldBindJSON(&at); err != nil {
	// 	restErr := errors.NewBadRequestError("invalid json body")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }

	// if err := handler.service.Create(&at); err != nil {
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	// c.JSON(http.StatusOK, at)

}
