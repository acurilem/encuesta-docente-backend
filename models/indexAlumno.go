package models

type IndexAlumno struct {
	AsignaturasInscritas  []SeccionesInscritasAlumno `json:"asignaturasInscritas"`
	PlanAlumno            []PlanAlumnoAno            `json:"planAlumno"`
	Contestadas           []Contestada               `json:"contestadas"`
	AsistenciaAsignaturas []AsistenciaAsign          `json:"asistenciaAsignaturas"`
	ProfesorSeleccionado  []ProfesorSeleccionado     `json:"profesorSeleccionado"`
	Profesores            []ProfesoresResumen        `json:"profesores"`
	InfoPlan              []PlanAlumnoActual         `json:"infoPlan"`
	AnoProcesoNum         int                        `json:"anoProcesoNum"`
	AnoIngreso            int                        `json:"anoIngreso"`
	NivelContesta         int                        `json:"nivelContesta"`
}
