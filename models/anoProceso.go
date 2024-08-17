package models

type AnoProceso struct {
	AnoProc string `gorm:"column:ano_proc" json:"anoProc"`
	SemProc string `gorm:"column:sem_proc" json:"semProc"`
}

// TableName overrides the table name used by User to `profiles`
func (AnoProceso) TableName() string {
	return "ano_proceso"
}
