package routes

import (
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/controller"
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/middleware"

	//"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitEncuestaDocenteRespuestasRoutes(r *gin.RouterGroup) {

	EncuestaDocenteRespuestasGroup := r.Group("/encuestaDocenteRespuestas")
	{
		//Grupo de rutas
		EncuestaDocenteRespuestasGroup.Use(middleware.SetRoles("ALUMNOS"), middleware.AuthMiddleware())
		{
			EncuestaDocenteRespuestasGroup.POST("/", controller.CreateEncuestaDocenteRespuestas)
			EncuestaDocenteRespuestasGroup.GET("/indexAlumno", controller.GetIndexAlumno)
			EncuestaDocenteRespuestasGroup.GET("/porContestar", controller.GetConfirmacionEncuestasPorContestarRespuestas)
		}
	}

}
