package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yimialmonte/src/domain/access_token"
	"github.com/yimialmonte/src/repository/db"
	"github.com/yimialmonte/src/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	service := access_token.NewService(db.New())
	handler := rest.NewHandler(service)

	router.GET("/oauth/access_token/:access_token_id", handler.GetById)

	router.Run(":8080")
}
