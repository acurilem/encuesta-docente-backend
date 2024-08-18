// main.go

package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/docs"
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/middleware"
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/routes"
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

//	@title			Encuesta docente
//	@version		1.0
//	@description	Módulo que incluye el apartado para la selección, desarrollo y envío de una encuesta docente de un alumno a un profesor.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	CITIAPS
//	@contact.url	https://citiaps.cl
//	@contact.email	citiaps@usach.cl

// lincense.name  Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				BearerJWTToken in Authorization Header

//	@accept		json
//	@produce	json

// @schemes	http https
func main() {
	utils.LoadEnv()
	gin.SetMode(os.Getenv("GIN_MODE"))
	app := gin.Default()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, log.Writer())
	app.Use(middleware.CorsMiddleware())
	// Docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Host = "localhost:8000"
	// Route docs
	app.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Servicio no encontrado."})
	})
	routes.InitRoutes(app)
	http.ListenAndServe(os.Getenv("ADDR"), app)
}
