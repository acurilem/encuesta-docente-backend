package models

type EncuestaDocenteContestadas struct {
	ID                   uint `gorm:"column:id" gorm:"primarykey"  json:"id"`
	IdCargaTotalProfesor int  `gorm:"column:id_carga_total_profesor" json:"idCargaTotalProfesor"`
	Ano                  int  `gorm:"column:ano" json:"ano"`
	Sem                  int  `gorm:"column:sem" json:"sem"`
	CodPersona           int  `gorm:"column:cod_alumno" json:"codPersona"` //cod_alumno es ub codPersona
}

// TableName overrides the table name
func (EncuestaDocenteContestadas) TableName() string {
	return "enc_docente_contestadas"
}
