package models

type PorcentajeAprobacionPlanCerti struct {
	RamosAprobados int `gorm:"column:ramos_aprobados" json:"ramosAprobados"`
	RamosMalla     int `gorm:"column:ramos_malla" json:"ramosMalla"`
	Porce          int `gorm:"column:porce" json:"porce"`
}
