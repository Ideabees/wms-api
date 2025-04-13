package routes

import (
	"wms-app/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	return r
}
