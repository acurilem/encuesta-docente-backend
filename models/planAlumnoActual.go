package models

type PlanAlumnoActual struct {
	CodTipoMatricula int    `gorm:"column:cod_tipo_matricula" json:"codTipoMatricula"`
	CodEspecialidad  int    `gorm:"column:cod_especialidad" json:"codEspecialidad"`
	CodMencion       int    `gorm:"column:cod_mencion" json:"codMencion"`
	CodCarrera       string `gorm:"column:cod_carrera" json:"codCarrera"`
	IdMatricula      int    `gorm:"column:id_matricula" json:"idMatricula"`
	AnoMatricula     int    `gorm:"column:ano_matricula" json:"anoMatricula"`
	SemEstadoFinal   int    `gorm:"column:sem_estado_final" json:"semEstadoFinal"`
	AnoProceso       int    `gorm:"column:ano_proceso" json:"anoProceso"`
	CodEstadoFinal   int    `gorm:"column:cod_estado_final" json:"codEstadoFinal"`
	IdPlanes         int    `gorm:"column:id_planes" json:"idPlanes"`
	NomCarrera       string `gorm:"column:nom_carrera" json:"nomCarrera"`
	NomMencion       string `gorm:"column:nom_mencion" json:"nomMencion"`
	AnoIngreso       int    `gorm:"column:ano_ingreso" json:"anoIngreso"`
	Anoplan          int    `gorm:"column:anoplan" json:"anoplan"`
	Descripcion      string `gorm:"column:descripcion" json:"descripcion"`
	NomEspecialidad  string `gorm:"column:nom_especialidad" json:"nomEspecialidad"`
	CodTipoPlan      int    `gorm:"column:cod_tipo_plan" json:"codTipoPlan"`
	CodNivelGlobal   int    `gorm:"column:cod_nivel_global" json:"codNivelGlobal"`
	NomCentro        string `gorm:"column:nom_centro" json:"nomCentro"`
	CodNivelCarrera  int    `gorm:"column:cod_nivel_carrera" json:"codNivelCarrera"`
}
