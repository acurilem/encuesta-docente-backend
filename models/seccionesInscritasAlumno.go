package models

import (
	"time"
)

type SeccionesInscritasAlumno struct {
	IdUsuarioAula        int       `gorm:"column:id_usuario_aula" json:"idUsuarioAula"`
	CodPersona           int       `gorm:"column:cod_persona" json:"codPersona"`
	IdCargaTotal         int       `gorm:"column:id_carga_total" json:"idCargaTotal"`
	Sem                  int       `gorm:"column:sem" json:"sem"`
	Ano                  int       `gorm:"column:ano" json:"ano"`
	IdTipoUsuario        int       `gorm:"column:id_tipo_usuario" json:"idTipoUsuario"`
	CodEstadoInscripcion int       `gorm:"column:cod_estado_inscripcion" json:"codEstadoInscripcion"`
	Opcion               int       `gorm:"column:opcion" json:"opcion"`
	IdMalla              int       `gorm:"column:id_malla" json:"idMalla"`
	IdMallaCd            int       `gorm:"column:id_malla_cd" json:"idMallaCd"`
	IdCargaParcial       int       `gorm:"column:id_carga_parcial" json:"idCargaParcial"`
	NomGrupo             string    `gorm:"column:nom_grupo" json:"nomGrupo"`
	CantidadEstimadaAlum int       `gorm:"column:cantidad_estimada_alum" json:"cantidadEstimadaAlum"`
	FechaIngreso         time.Time `gorm:"column:fecha_ingreso" json:"fechaIngreso"`
	Login                string    `gorm:"column:login" json:"login"`
	CodTipoSeccion       int       `gorm:"column:cod_tipo_seccion" json:"codTipoSeccion"`
	NumGrupos            int       `gorm:"column:num_grupos" json:"numGrupos"`
	HCP                  float64   `gorm:"column:HCP" json:"HCP"`
	IdDetalleCategoria   int       `gorm:"column:id_detalle_categoria" json:"idDetalleCategoria"`
	CodCategoria         int       `gorm:"column:cod_categoria" json:"codCategoria"`
	Ambito               int       `gorm:"column:ambito" json:"ambito"`
}
