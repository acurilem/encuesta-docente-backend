package models

import "time"

type MatriculaAlumnoAnoIngresoAnoproceso struct {
	IdMatricula        int       `gorm:"column:id_matricula" json:"idMatricula"`
	IdHojaRuta         int       `gorm:"column:id_hoja_ruta" json:"idHojaRuta"`
	CodPersona         int       `gorm:"column:cod_persona" json:"codPersona"`
	RutAlum            string    `gorm:"column:rut_alum" json:"rutAlum"`
	AnoProceso         int       `gorm:"column:ano_proceso" json:"anoProceso"`
	AnoMatricula       int       `gorm:"column:ano_matricula" json:"anoMatricula"`
	CodTipoAlumno      int       `gorm:"column:cod_tipo_alumno" json:"codTipoAlumno"`
	CodModoIngreso     int       `gorm:"column:cod_modo_ingreso" json:"codModoIngreso"`
	AnoIngreso         int       `gorm:"column:ano_ingreso" json:"anoIngreso"`
	IdPlanes           int       `gorm:"column:id_planes" json:"idPlanes"`
	CodTipoMatricula1  int       `gorm:"column:cod_tipo_matricula" json:"codTipoMatricula1"`
	IdMatriculaColegio int       `gorm:"column:id_matricula_colegio" json:"idMatriculaColegio"`
	FechaMatricula     time.Time `gorm:"column:fecha_matricula" json:"fechaMatricula"`
	CodEstadoRetiro    int       `gorm:"column:cod_estado_retiro" json:"codEstadoRetiro"`
	CodEstadoFinal     int       `gorm:"column:cod_estado_final" json:"codEstadoFinal"`
	SemEstadoFinal     int       `gorm:"column:sem_estado_final" json:"semEstadoFinal"`
	CodTipoAlumMatr    int       `gorm:"column:cod_tipo_alum_matr" json:"codTipoAlumMatr"`
	CodTipoPrueba      int       `gorm:"column:cod_tipo_prueba" json:"codTipoPrueba"`
	AnoPrimerIngreso   int       `gorm:"column:ano_primer_ingreso" json:"anoPrimerIngreso"`
	SemDesvinculacion  int       `gorm:"column:sem_desvinculacion" json:"semDesvinculacion"`
	IdPrueba           int       `gorm:"column:id_prueba" json:"idPrueba"`
	Posicion           int       `gorm:"column:posicion" json:"posicion"`
	IdFichaTesoreria   int       `gorm:"column:id_ficha_tesoreria" json:"idFichaTesoreria"`
	Descripcion1       string    `gorm:"column:descripcion" json:"descripcion1"`
	CodTipoMatricula2  int       `gorm:"column:cod_tipo_matricula" json:"codTipoMatricula2"`
	Descripcion2       string    `gorm:"column:descripcion" json:"descripcion2"`
	FechaMatricula1    string    `gorm:"column:fecha_matricula1" json:"fechaMatricula1"`
	TipoMatricula      string    `gorm:"column:tipo_matricula" json:"tipoMatricula"`
}
