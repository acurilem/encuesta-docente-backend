package models

type ProfesoresResumen struct {
	IdCargaTotal       int                  `json:"idCargaTotalProfesor"`
	CargaTotalProfesor []CargaTotalProfesor `json:"cargaTotalProfesor"`
}
