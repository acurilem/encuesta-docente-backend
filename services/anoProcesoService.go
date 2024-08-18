package services

import (
	"errors"
	"reflect"
	"time"

	"github.com/acurilem/encuesta-docente-backend/config"
	"github.com/acurilem/encuesta-docente-backend/models"
)

// Getall de un Año proceso
// Ano_proceso_actual()
func GetAllAnoProcesoService() ([]models.AnoProceso, error) {
	// Se establece conexión
	db := config.Database

	err := db.Error
	if err != nil {
		return nil, err
	}

	// Variable que contiene a todos los AnoProceso
	var anoProceso []models.AnoProceso
	// Trae a todos los AnoProceso desde la base de datos
	result := db.Find(&anoProceso)
	err = result.Error
	if err != nil {
		return nil, err
	}
	//En caso de algun error
	if result.Error != nil {
		return anoProceso, result.Error
	}
	return anoProceso, nil
}

// Getall de un matricula Ano Proceso
// Ano_proceso_matricula()
func GetAllMatriculaAnoProcesoService() ([]models.MatriculaAnoProceso, error) {
	// Se establece conexión
	db := config.Database
	// Variable que contiene a todos los matriculaAnoProceso
	var matriculaAnoProceso []models.MatriculaAnoProceso
	// Trae a todos los matriculaAnoProceso desde la base de datos
	result := db.Find(&matriculaAnoProceso)
	//En caso de algun error
	if result.Error != nil {
		return matriculaAnoProceso, result.Error
	}
	return matriculaAnoProceso, nil
}

// funcion para modificar un AnoProceso
// modificar_ano_proceso()
func UpdateAnoProcesoService(updatedAnoProceso models.AnoProceso, anoProcesoActual int) (models.AnoProceso, error) {
	// Se establece conexión
	db := config.Database

	// Se recorre el modelo updatedAnoProceso
	modelType := reflect.TypeOf(updatedAnoProceso)
	modelValue := reflect.ValueOf(updatedAnoProceso)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los ano_proc coincidan
		if value != 0 && value != "" {
			result := db.Model(&models.AnoProceso{}).Where("ano_proc=?", anoProcesoActual).Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedAnoProceso, result.Error
			} else if result.RowsAffected < 1 {
				return updatedAnoProceso, errors.New("AnoProceso no encontrado")
			}
		}
	}
	return updatedAnoProceso, nil
}

// Función para obtener un AdminPostUmagPsu por rut inscrito
// select_admin_post_umag_psu()
func GetAdminPostUmagPsuByRutService(rutInscrito string) (models.AdminPostUmagPsu, error) {
	// Se establece conexión
	db := config.Database
	// a partir del rut inscrito del AdminPostUmagPsu.
	var adminPostUmagPsu models.AdminPostUmagPsu
	// Obtiene el AdminPostUmagPsu.
	result := db.First(&adminPostUmagPsu, rutInscrito)
	//En caso de algun error
	if result.Error != nil {
		return adminPostUmagPsu, result.Error
	}
	// Devuelve el AdminPostUmagPsu encontrado
	return adminPostUmagPsu, nil

}

// Función para obtener un AdminPostUmagPsu por rutInscrito y año de proceso
// select_admin_post_umag_psu_por_ano()
func GetAdminPostUmagPsuByRutAnoProcService(rutInscrito string) ([]models.AdminPostUmagPsu, error) {
	// Se establece conexión
	db := config.Database

	var adminPostUmagPsu []models.AdminPostUmagPsu
	// Obtiene el año actual
	anoProc := time.Now().Year()
	// Obtiene el AdminPostUmagPsu.
	result := db.Where("rut_ins = ? AND ano_proc = ?", rutInscrito, anoProc).Find(&adminPostUmagPsu)

	//En caso de algun error
	if result.Error != nil {
		return adminPostUmagPsu, result.Error
	}
	// Devuelve el AdminPostUmagPsu encontrado
	return adminPostUmagPsu, nil

}

// funcion para modificar un AdminPostUmagPsu
// update_admin_post_umag_psu()
func UpdateAdminPostUmagPsuService(updatedAdminPostUmagPsu models.AdminPostUmagPsu, anoProcesoActual int, rutInscrito string) (models.AdminPostUmagPsu, error) {
	// Se establece conexión
	db := config.Database

	// Se recorre el modelo updatedAdminPostUmagPsu
	modelType := reflect.TypeOf(updatedAdminPostUmagPsu)
	modelValue := reflect.ValueOf(updatedAdminPostUmagPsu)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los ano_proc coincidan
		if value != 0 && value != "" {
			result := db.Model(&models.AdminPostUmagPsu{}).Where("rut_ins = ? AND ano_proc= ?", rutInscrito, anoProcesoActual).Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedAdminPostUmagPsu, result.Error
			} else if result.RowsAffected < 1 {
				return updatedAdminPostUmagPsu, errors.New("AdminPostUmagPsu no encontrado")
			}
		}
	}
	return updatedAdminPostUmagPsu, nil
}
