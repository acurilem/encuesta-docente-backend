package models

type EncuestaDocentePorContestar struct {
	IdCargaTotal   int `gorm:"column:id_carga_total" json:"idCargaTotal"`
	IdUsuarioAula  int `gorm:"column:id_usuario_aula" json:"idUsuarioAula"`
}
