package router

import (
	"goGeradorCnab/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	router.GET("/cnab", controller.GetCnab)

	router.Run("localhost:8080")
}
