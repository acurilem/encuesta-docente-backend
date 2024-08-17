package services

import (
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/config"
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/models"
)

// Función para crear un RastroSayd en la base de datos
func CreateRastroSaydService(newRastroSayd models.RastroSayd) (models.RastroSayd, error) {
	// Se establece conexión
	db := config.Database

	// Inserta el RastroSayd en la colección.
	result := db.Create(&newRastroSayd)
	//En caso de algun error
	if result.Error != nil {
		return newRastroSayd, result.Error
	}
	return newRastroSayd, nil
}
