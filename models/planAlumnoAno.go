package models

import "time"

type PlanAlumnoAno struct {
	IdMatricula        int       `json:"idMatricula"`
	IdHojaRuta         int       `json:"idHojaRuta"`
	CodPersona         int       `json:"codPersona"`
	RutAlum            string    `json:"rutAlum"`
	AnoProceso         int       `json:"anoProceso"`
	AnoMatricula       int       `json:"anoMatricula"`
	CodTipoAlumno      int       `json:"codTipoAlumno"`
	CodModoIngreso     int       `json:"codModoIngreso"`
	AnoIngreso         int       `json:"anoIngreso"`
	IdPlanes           int       `json:"idPlanes"`
	CodTipoMatricula   int       `json:"codTipoMatricula"`
	IdMatriculaColegio int       `json:"idMatriculaColegio"`
	FechaMatricula     time.Time `json:"fechaMatricula"`
	CodEstadoRetiro    int       `json:"codEstadoRetiro"`
	CodEstadoFinal     int       `json:"codEstadoFinal"`
	SemEstadoFinal     int       `json:"semEstadoFinal"`
	CodTipoAlumMatr    int       `json:"codTipoAlumMatr"`
	CodTipoPrueba      int       `json:"codTipoPrueba"`
	AnoPrimerIngreso   int       `json:"anoPrimerIngreso"`
	SemDesvinculacion  int       `json:"semDesvinculacion"`
	IdPrueba           int       `json:"idPrueba"`
	Posicion           int       `json:"posicion"`
	IdFichaTesoreria   int       `json:"idFichaTesoreria"`
	IdCarreraTotal     int       `json:"idCarreraTotal"`
	Anoplan            int       `json:"anoplan"`
	Vigencia           int       `json:"vigencia"`
	CodTipoPlan        int       `json:"codTipoPlan"`
	Duracion           int       `json:"duracion"`
	CodRegimen         int       `json:"codRegimen"`
	CodNivelGlobal     int       `json:"codNivelGlobal"`
	CodNivelCarrera    int       `json:"codNivelCarrera"`
	Caracteristica     string    `json:"caracteristica"`
	CodJornada         int       `json:"codJornada"`
	CodModalidad       int       `json:"codModalidad"`
}
