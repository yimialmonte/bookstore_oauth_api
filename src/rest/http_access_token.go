package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yimialmonte/src/domain/access_token"
)

type AccesTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(s access_token.Service) AccesTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (a *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := a.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
