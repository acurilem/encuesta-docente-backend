package middleware

import (
	"io/ioutil"
	"net/http"
	"os"

	// "sid3-api/services"
	"time"

	"github.com/acurilem/SID-UMAG-Encuesta-docente-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tidwall/gjson"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))
var backendURL = os.Getenv("ROL_SERVICE")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := services.GetToken(c)
		token, err := jwt.ParseWithClaims(tokenString, &services.Sid3Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		}, jwt.WithLeeway(5*time.Second))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error de autenticación"})
			return
		}

		if claims, ok := token.Claims.(*services.Sid3Claims); ok && token.Valid {
			if claims.CodPersona > 0 {
				roles, exists := c.Get("roles")
				if exists {

					rolesAuth, ok := roles.([]string)

					if !ok {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener roles"})
						c.Abort()
						return
					}

					// Solicitud Get
					request, err := http.NewRequest("GET", backendURL, nil)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
						c.Abort()
						return
					}

					request.Header.Set("Authorization", "Bearer "+tokenString)

					response, err := http.DefaultClient.Do(request)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Request failed"})
						c.Abort()
						return
					}

					defer response.Body.Close()

					// Respuesta
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "No response"})
						c.Abort()
						return
					}

					roles := gjson.Get(string(body), "Roles").Array()
					for _, role := range roles {
						for _, roleAuth := range rolesAuth {
							if role.Get("descripcion").String() == roleAuth {
								c.Next()
								return
							}
						}
					}
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Rol no autorizado"})
					return

				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no valido"})
				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

	}
}

func SetRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("roles", roles)
		c.Next()
	}
}
