package routes

import (
	"github.com/acurilem/encuesta-docente-backend/controller"
	"github.com/acurilem/encuesta-docente-backend/middleware"

	//"github.com/acurilem/encuesta-docente-backend/middleware"
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
