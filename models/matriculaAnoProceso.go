package models

type MatriculaAnoProceso struct {
	AnoProc        uint `gorm:"column:ano_proc"  json:"anoProc"`
	AnoInscripcion uint `gorm:"column:ano_inscripcion" json:"anoInscripcion "`
}

// TableName overrides the table name used by User to `profiles`
func (MatriculaAnoProceso) TableName() string {
	return "matricula_ano_proceso"
}
