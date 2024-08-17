package models

type EncuestaDocenteRespuestas struct {
	ID                   uint   `gorm:"column:id" gorm:"primarykey"  json:"id"`
	IdCargaTotalProfesor int    `gorm:"column:id_carga_total_profesor" json:"idCargaTotalProfesor"`
	P1                   int    `gorm:"column:p1" json:"p1"`
	P2                   int    `gorm:"column:p2" json:"p2"`
	P3                   int    `gorm:"column:p3" json:"p3"`
	P4                   int    `gorm:"column:p4" json:"p4"`
	P5                   int    `gorm:"column:p5" json:"p5"`
	P6                   int    `gorm:"column:p6" json:"p6"`
	P7                   int    `gorm:"column:p7" json:"p7"`
	P8                   int    `gorm:"column:p8" json:"p8"`
	Nota                 int    `gorm:"column:nota" json:"nota"`
	Comentario           string `gorm:"column:comentario" json:"comentario"`
	Ano                  int    `gorm:"column:ano" json:"ano"`
	Sem                  int    `gorm:"column:sem" json:"sem"`
}

// TableName overrides the table name
func (EncuestaDocenteRespuestas) TableName() string {
	return "enc_docente_respuestas"
}
