package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swjang1214/bookstore_oauth-api/src/domain/access_token"
)

type IAccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.IService
}

func NewHandler(service access_token.IService) IAccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
