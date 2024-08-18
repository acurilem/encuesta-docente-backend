package controller

import (
	"log"
	"net/http"

	"github.com/acurilem/encuesta-docente-backend/services"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	user, err := services.GetUser(c)
	if err != nil {
		log.Println("No fue posible encontrar al usuario en SAYD / Vista Personas")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sin info en SAYD"})
		return
	}
	c.JSON(http.StatusOK, user)
}
