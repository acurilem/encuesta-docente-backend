package controller

import (
	"log"
	"net/http"

	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/models"
	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/services"

	"github.com/gin-gonic/gin"
)

func LoginFunc(c *gin.Context) {
	var loginValues models.Login
	if err := c.ShouldBindJSON(&loginValues); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Datos enviados no cunplen con módelo"})
		return
	}
	var rut string
	var err error
	if services.IsRutOrPassport(loginValues.Username) {
		rut, err = services.LoginLdapv2WithRut(loginValues)
	} else {
		rut, err = services.LoginLdapv2(loginValues)
	}

	if err != nil {
		log.Println("No fue posible encontrar al usuario en LDAP")
		c.AbortWithError(http.StatusUnauthorized, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas LDAP"})
		return
	}
	user, err := services.GetUserInfoFromRutService(rut)
	if err != nil {
		log.Println("No fue posible encontrar al usuario en SAYD / Vista Personas")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Sin info en SAYD"})
		return
	}
	token, refreshToken, err1, err2 := services.LoadJWTAuth(user.NombreCompleto, user.MailInstitucional, user.ID)
	if err1 != nil || err2 != nil {
		//if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func RefreshToken(c *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	var tokenRB tokenReqBody
	if err := c.ShouldBindJSON(&tokenRB); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	claims, err := services.ValidateRefreshToken(tokenRB.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de recuperación inválido"})
		c.Abort()
		return
	}
	user, err := services.GetUserInfoFromCodPersonaService(uint(claims.RefreshCodPersona))
	token, refreshToken, err1, err2 := services.LoadJWTAuth(user.NombreCompleto, user.MailInstitucional, user.ID)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo regenerar el token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
		"user":         user,
	})
}
