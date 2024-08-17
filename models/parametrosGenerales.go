package models

import "time"

type ParametrosGenerales struct {
	CodTipoParametro     int       `gorm:"column:cod_tipo_parametro" json:"codTipoParametro"`
	DescripcionParametro string    `gorm:"column:descripcion_parametro" json:"descripcionParametro"`
	Switch               int       `gorm:"column:switch" json:"switch"`
	Fecha                time.Time `gorm:"column:fecha" json:"fecha"`
	Condicion1           string    `gorm:"column:condicion1" json:"condicion1"`
	Condicion2           string    `gorm:"column:condicion2" json:"condicion2"`
	Observacion          string    `gorm:"column:observacion" json:"observacion"`
}

func (ParametrosGenerales) TableName() string {
	return "parametros_generales"
}
