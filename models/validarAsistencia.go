package models

type ValidarAsistencia struct {
	IdAsistencia         int `gorm:"column:id_asistencia" json:"idAsistencia"`
	IdUsuarioAula        int `gorm:"column:id_usuario_aula" json:"idUsuarioAula"`
	Porcentaje           int `gorm:"column:porcentaje" json:"porcentaje"`
	CodPersona           int `gorm:"column:cod_persona" json:"codPersona"`
	IdCargaTotal         int `gorm:"column:id_carga_total" json:"idCargaTotal"`
	Sem                  int `gorm:"column:sem" json:"sem"`
	Ano                  int `gorm:"column:ano" json:"ano"`
	IdTipoUsuario        int `gorm:"column:id_tipo_usuario" json:"idTipoUsuario"`
	CodEstadoInscripcion int `gorm:"column:cod_estado_inscripcion" json:"codEstadoInscripcion"`
	Opcion               int `gorm:"column:opcion" json:"opcion"`
	IdMalla              int `gorm:"column:id_malla" json:"idMalla"`
	IdMallaCd            int `gorm:"column:id_malla_cd" json:"idMallaCd"`
}
