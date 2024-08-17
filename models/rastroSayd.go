package models

import (
	"time"
)

type RastroSayd struct {
	Login      string    `gorm:"column:login"  json:"login"`
	Fecha      time.Time `gorm:"column:fecha"  json:"fecha"`
	Operacion  string    `gorm:"column:operacion"  json:"operacion"`
	Aplicacion string    `gorm:"column:aplicacion "  json:"aplicacion"`
	Datos      string    `gorm:"column:datos"  json:"datos"`
	Tabla      string    `gorm:"column:tabla"  json:"tabla"`
	Version    string    `gorm:"column:version"  json:"version"`
}

func (RastroSayd) TableName() string {
	return "rastro_sayd"
}
