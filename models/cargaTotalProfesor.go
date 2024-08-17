package models

type CargaTotalProfesor struct {
	IdCargaTotalProfesor int    `gorm:"column:id_carga_total_profesor" json:"idCargaTotalProfesor"`
	NombreCompleto       string `gorm:"column:nombre_completo" json:"nombreCompleto"`
}
