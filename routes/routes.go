package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	groupRouter := r.Group("/api/v1/encuestaDocenteRespuestas")

	InitEncuestaDocenteRespuestasRoutes(groupRouter)
	InitAuthRoutes(groupRouter)

}
