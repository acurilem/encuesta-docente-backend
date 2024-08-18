package routes

import (
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/controller"
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitAuthRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controller.LoginFunc)
		authGroup.POST("/refresh", controller.RefreshToken)
		authenticatedRoutes := r.Group("/auth")
		authenticatedRoutes.Use(middleware.AuthMiddleware())
		{
			authenticatedRoutes.GET("/user", controller.User)
		}
	}
}
