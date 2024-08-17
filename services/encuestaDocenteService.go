package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/config"
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/models"
)

// Guarda las respuestas de la encuesta docente de un alumno y deja rastro de ella.
// guardar_respuestas_encuesta_docente(
func CreateEncuestaDocenteRespuestasService(codPersona int, respuestas models.EncuestaDocenteRespuestas) (models.EncuestaDocenteRespuestas, error) {
	db := config.Database
	//guarda las respuestas de encuesta docente
	result := db.Create(&respuestas)
	//En caso de algun error
	if result.Error != nil {
		return respuestas, result.Error
	}

	// Se obtiene los datos del usuario, lo que se hace es obtener su correo y hacer un split por el @
	user, err := GetUserInfoFromCodPersonaService(uint(codPersona))
	if err != nil {
		return respuestas, err
	}
	//a partir del correo se obtiene el username
	partes := strings.Split(user.MailSid, "@")
	usuario := partes[0]
	// Se deja rastro del insert realizado
	// Para ello se crea un rastro
	var rastro models.RastroSayd
	rastro.Login = usuario
	rastro.Fecha = time.Now()
	rastro.Operacion = "INSERT"
	rastro.Aplicacion = "encuesta evaluacion docente"
	rastro.Datos = "INSERT INTO enc_docente_respuestas VALUES(" + strconv.Itoa(respuestas.IdCargaTotalProfesor) + ", " + strconv.Itoa(respuestas.P1) + ", " + strconv.Itoa(respuestas.P2) + ", " + strconv.Itoa(respuestas.P3) + ", " + strconv.Itoa(respuestas.P4) + ", " + strconv.Itoa(respuestas.P5) + ", " + strconv.Itoa(respuestas.P6) + ", " + strconv.Itoa(respuestas.P7) + ", " + strconv.Itoa(respuestas.P8) + ", " + strconv.Itoa(respuestas.Nota) + ", " + respuestas.Comentario + ", " + strconv.Itoa(respuestas.Ano) + ", " + strconv.Itoa(respuestas.Sem) + ");"
	rastro.Tabla = "enc_docente_respuestas"
	createdRastro, err := CreateRastroSaydService(rastro)
	// En caso de algun error
	if err != nil {
		createdRastro.Aplicacion = "Fallo"
		return respuestas, result.Error
	}

	var encuestaCont models.EncuestaDocenteContestadas
	encuestaCont.IdCargaTotalProfesor = respuestas.IdCargaTotalProfesor
	encuestaCont.Sem = respuestas.Sem
	encuestaCont.Ano = respuestas.Ano
	encuestaCont.CodPersona = codPersona
	//deja registro de que la persona contesto la encuesta
	result = db.Create(&encuestaCont)
	//En caso de algun error
	if result.Error != nil {
		return respuestas, result.Error
	}

	// Se deja rastro del insert realizado
	// Para ello se crea un rastro
	var rastro2 models.RastroSayd
	rastro2.Login = usuario
	rastro2.Fecha = time.Now()
	rastro2.Operacion = "INSERT"
	rastro2.Aplicacion = "encuesta evaluacion docente"
	rastro2.Datos = "INSERT INTO enc_docente_contestadas VALUES(" + strconv.Itoa(respuestas.IdCargaTotalProfesor) + ", " + strconv.Itoa(codPersona) + ", " + strconv.Itoa(respuestas.Ano) + ", " + strconv.Itoa(respuestas.Sem) + ");"
	rastro2.Tabla = "enc_docente_contestadas"
	createdRastro2, err := CreateRastroSaydService(rastro2)
	// En caso de algun error
	if err != nil {
		createdRastro2.Aplicacion = "Fallo"
		return respuestas, result.Error
	}

	return respuestas, nil
}

// Seleciona las secciones  inscritas de un alumno
// secciones_inscritas_alumno()
func GetSeccionesInscritasAlumnoService(ano string, sem string, codPersona int) ([]models.SeccionesInscritasAlumno, error) {
	db := config.Database
	var results []models.SeccionesInscritasAlumno
	codPersonaAux := strconv.Itoa(codPersona)

	cond1 := ""
	if sem == "2" {
		cond1 = "UNION SELECT ua.*, ct.*, dc.* FROM usuarios_aula ua INNER JOIN malla_pionero mp ON mp.id_malla = ua.id_malla INNER JOIN carga_total ct ON ua.id_carga_total = ct.id_carga_total INNER JOIN detalle_categoria dc ON ua.cod_persona = dc.cod_persona WHERE (ua.ano = " + ano + ") AND (ua.sem = '1') AND (dc.cod_categoria = 40) AND (ua.cod_persona =" + codPersonaAux + ") AND ua.cod_estado_inscripcion IN (1, 2, 10, 12 /*, 14*/) AND mp.duracion = 'A'"
	}
	err := db.Raw("SELECT ua.*,ct.*, dc.* FROM usuarios_aula ua INNER JOIN carga_total ct ON ua.id_carga_total = ct.id_carga_total INNER JOIN detalle_categoria dc ON ua.cod_persona = dc.cod_persona WHERE (ua.ano = ?) AND (ua.sem = ?) AND (dc.cod_categoria = 40) AND (ua.cod_persona = ?) AND ua.cod_estado_inscripcion IN (1,2,10,12/*,14*/)"+cond1+" ORDER BY nom_grupo", ano, sem, codPersona).Scan(&results).Error
	return results, err
}

// busca el plan de un alumno para un año especifico
// plan_alumno_ano(
func GetPlanAlumnoAnoService(codPersona int, ano int) ([]models.PlanAlumnoAno, error) {
	db := config.Database
	var results []models.PlanAlumnoAno
	err := db.Raw("SELECT * FROM matricula INNER JOIN planes ON matricula.id_planes = planes.id_planes WHERE (matricula.cod_persona = ?) AND (matricula.ano_proceso = ?)", codPersona, ano).Scan(&results).Error

	return results, err
}

// valida la asistencia de un alumno
// valida_asistencia(
func GetValidarAsistenciaService(idUsuarioAula int) ([]models.ValidarAsistencia, error) {
	db := config.Database
	var results []models.ValidarAsistencia
	err := db.Raw("SELECT * FROM asistencia_alumnos aa INNER JOIN usuarios_aula ua ON aa.id_usuario_aula = ua.id_usuario_aula WHERE (ua.id_usuario_aula = ?);", idUsuarioAula).Scan(&results).Error

	return results, err
}

// Obtiene el id de la carga total de un profesor y su nombre
// profesores_carga_total(
func GetProfesoresCargaTotalService(idCargaTotal int) ([]models.CargaTotalProfesor, error) {
	db := config.Database
	var results []models.CargaTotalProfesor
	err := db.Raw("SELECT MIN(ctp.id_carga_total_profesor) AS id_carga_total_profesor, vp.nombre_completo FROM carga_total_profesores ctp INNER JOIN vista_personas vp ON ctp.cod_persona=vp.cod_persona WHERE ctp.id_carga_total=? AND vp.cod_persona!='30283' GROUP BY vp.nombre_completo ORDER BY vp.nombre_completo", idCargaTotal).Scan(&results).Error

	return results, err
}

// valida si el alumno ya contesto una encuesta de alguna sección
// valida_contesto_encuesta_alguna_seccion(
func GetValidarEncuestasContestadasService(idCargaTotal int, ano int, sem string, codPersona int) (int, error) {
	db := config.Database
	var results []map[string]interface{}
	if sem == "2" {
		err := db.Raw("select * from (SELECT mp.id_malla, cd.id_curso_dictar, cap.id_carga_parcial, cap.tipo_curso, cap.cant_grupos_teoria AS cant_grupo_teo, cap.cant_grupos_practica AS cant_grupo_pra,ct.cod_tipo_seccion AS tipo_curso_seccion, ct.nom_grupo, ca.id_comp_asig, ca.nom_comp_asig, ct.id_carga_total, ca.nom_comp_asig AS nom_comun, cp.sct_total, mp.nro_sem, mp.nro_ano, mp.cod_tipo_curso_malla, cap.hrs_crono_presen_semestral AS hr_total, cap.hrs_teoria AS hr_teoria, cap.hrs_practica AS hr_practica, ct.HCP AS hr_real, cd.ano, cd.sem, cp.hrs_teoricas AS t, cp.hrs_ejercicios AS e, cp.hrs_practicas AS l, cp.es_TEL AS es_tel, mp.id_planes, mp.duracion, cd.se_dicta, ct.num_grupos, ct.cantidad_estimada_alum,cd.cod_unidad_responsable FROM sayd.malla_pionero AS mp INNER JOIN sayd.cursos_dictar AS cd ON cd.id_malla = mp.id_malla INNER JOIN sayd.carga_parcial AS cap ON cap.cod_curso_dictar = cd.id_curso_dictar INNER JOIN sayd.carga_total AS ct ON ct.id_carga_parcial = cap.id_carga_parcial INNER JOIN sayd.cursos_pionero AS cp ON cp.id_curso_pionero = mp.id_curso_pionero INNER JOIN sayd.competencia_asignatura AS ca ON ca.id_comp_asig = cp.id_comp_asig WHERE (cap.tipo_curso = 0) AND (cd.id_curso_dictar NOT IN (SELECT id_curso_dictar FROM sayd.detalle_curso_comun)) UNION ALL SELECT TOP 100 PERCENT mp.id_malla, cd.id_curso_dictar, cap.id_carga_parcial, cap.tipo_curso, cap.cant_grupos_teoria AS cant_grupo_teo, cap.cant_grupos_practica AS cant_grupo_pra, ct.cod_tipo_seccion AS tipo_curso_seccion, ct.nom_grupo, ca.id_comp_asig, ca.nom_comp_asig, ct.id_carga_total, cc.nom_curso_comun AS nom_comun, cp.sct_total, mp.nro_sem, mp.nro_ano, mp.cod_tipo_curso_malla, cap.hrs_crono_presen_semestral AS hr_total, cap.hrs_teoria AS hr_teoria, cap.hrs_practica AS hr_practica, ct.HCP AS hr_real, cd.ano, cd.sem, cp.hrs_teoricas AS t, cp.hrs_ejercicios AS e, cp.hrs_practicas AS l, cp.es_TEL AS es_tel, mp.id_planes, mp.duracion, cd.se_dicta, ct.num_grupos, ct.cantidad_estimada_alum,cd.cod_unidad_responsable FROM sayd.malla_pionero mp INNER JOIN sayd.cursos_dictar cd ON cd.id_malla = mp.id_malla INNER JOIN sayd.detalle_curso_comun dcc ON dcc.id_curso_dictar = cd.id_curso_dictar INNER JOIN sayd.curso_comun cc ON cc.id_curso_comun = dcc.id_curso_comun INNER JOIN sayd.carga_parcial cap ON cap.cod_curso_dictar = cc.id_curso_comun INNER JOIN sayd.carga_total ct ON ct.id_carga_parcial = cap.id_carga_parcial INNER JOIN sayd.cursos_pionero cp ON cp.id_curso_pionero = mp.id_curso_pionero INNER JOIN sayd.competencia_asignatura ca ON ca.id_comp_asig = cp.id_comp_asig WHERE (cap.tipo_curso = 1)) as a where a.id_carga_total=? and a.ano=? and a.sem=1 and a.duracion='A'", idCargaTotal, ano).Scan(&results).Error
		if err != nil {
			return -1, err
		}
		if len(results) > 0 { // Si hay al menos una fila, es una asignatura anual
			sem = "1"
		}
	}
	var results2 []map[string]interface{}
	err := db.Raw("select * from enc_docente_contestadas edc INNER JOIN carga_total_profesores ctp ON edc.id_carga_total_profesor=ctp.id_carga_total_profesor WHERE edc.cod_alumno=? AND ano=? AND sem=? AND ctp.id_carga_total=?", codPersona, ano, sem, idCargaTotal).Scan(&results2).Error
	if err != nil {
		return -1, err
	}
	if len(results2) > 0 {
		return 1, err
	}
	return 0, err
}

// retorna la informacion del profesor seleccionado
// profesor_seleccionado(
func GetProfesorSeleccionadoService(idCargaTotal int, ano int, sem string, codPersona int) (string, error) {
	db := config.Database
	var results []map[string]interface{}
	if sem == "2" {
		err := db.Raw("select * from (SELECT mp.id_malla, cd.id_curso_dictar, cap.id_carga_parcial, cap.tipo_curso, cap.cant_grupos_teoria AS cant_grupo_teo, cap.cant_grupos_practica AS cant_grupo_pra,ct.cod_tipo_seccion AS tipo_curso_seccion, ct.nom_grupo, ca.id_comp_asig, ca.nom_comp_asig, ct.id_carga_total, ca.nom_comp_asig AS nom_comun, cp.sct_total, mp.nro_sem, mp.nro_ano, mp.cod_tipo_curso_malla, cap.hrs_crono_presen_semestral AS hr_total, cap.hrs_teoria AS hr_teoria, cap.hrs_practica AS hr_practica, ct.HCP AS hr_real, cd.ano, cd.sem, cp.hrs_teoricas AS t, cp.hrs_ejercicios AS e, cp.hrs_practicas AS l, cp.es_TEL AS es_tel, mp.id_planes, mp.duracion, cd.se_dicta, ct.num_grupos, ct.cantidad_estimada_alum,cd.cod_unidad_responsable FROM sayd.malla_pionero AS mp INNER JOIN sayd.cursos_dictar AS cd ON cd.id_malla = mp.id_malla INNER JOIN sayd.carga_parcial AS cap ON cap.cod_curso_dictar = cd.id_curso_dictar INNER JOIN sayd.carga_total AS ct ON ct.id_carga_parcial = cap.id_carga_parcial INNER JOIN sayd.cursos_pionero AS cp ON cp.id_curso_pionero = mp.id_curso_pionero INNER JOIN sayd.competencia_asignatura AS ca ON ca.id_comp_asig = cp.id_comp_asig WHERE (cap.tipo_curso = 0) AND (cd.id_curso_dictar NOT IN (SELECT id_curso_dictar FROM sayd.detalle_curso_comun)) UNION ALL SELECT TOP 100 PERCENT mp.id_malla, cd.id_curso_dictar, cap.id_carga_parcial, cap.tipo_curso, cap.cant_grupos_teoria AS cant_grupo_teo, cap.cant_grupos_practica AS cant_grupo_pra, ct.cod_tipo_seccion AS tipo_curso_seccion, ct.nom_grupo, ca.id_comp_asig, ca.nom_comp_asig, ct.id_carga_total, cc.nom_curso_comun AS nom_comun, cp.sct_total, mp.nro_sem, mp.nro_ano, mp.cod_tipo_curso_malla, cap.hrs_crono_presen_semestral AS hr_total, cap.hrs_teoria AS hr_teoria, cap.hrs_practica AS hr_practica, ct.HCP AS hr_real, cd.ano, cd.sem, cp.hrs_teoricas AS t, cp.hrs_ejercicios AS e, cp.hrs_practicas AS l, cp.es_TEL AS es_tel, mp.id_planes, mp.duracion, cd.se_dicta, ct.num_grupos, ct.cantidad_estimada_alum,cd.cod_unidad_responsable FROM sayd.malla_pionero mp INNER JOIN sayd.cursos_dictar cd ON cd.id_malla = mp.id_malla INNER JOIN sayd.detalle_curso_comun dcc ON dcc.id_curso_dictar = cd.id_curso_dictar INNER JOIN sayd.curso_comun cc ON cc.id_curso_comun = dcc.id_curso_comun INNER JOIN sayd.carga_parcial cap ON cap.cod_curso_dictar = cc.id_curso_comun INNER JOIN sayd.carga_total ct ON ct.id_carga_parcial = cap.id_carga_parcial INNER JOIN sayd.cursos_pionero cp ON cp.id_curso_pionero = mp.id_curso_pionero INNER JOIN sayd.competencia_asignatura ca ON ca.id_comp_asig = cp.id_comp_asig WHERE (cap.tipo_curso = 1)) as a where a.id_carga_total=? and a.ano=? and a.sem=1 and a.duracion='A'", idCargaTotal, ano).Scan(&results).Error
		if err != nil {
			return "", err
		}
		if len(results) > 0 { // Si hay al menos una fila, es una asignatura anual
			sem = "1"
		}
	}
	var results2 ProfesorNombre
	err := db.Raw("select vp.nombre_completo from enc_docente_contestadas edc INNER JOIN carga_total_profesores ctp ON edc.id_carga_total_profesor=ctp.id_carga_total_profesor INNER JOIN vista_personas vp ON ctp.cod_persona=vp.cod_persona INNER JOIN carga_total ct ON ct.id_carga_total=ctp.id_carga_total where edc.cod_alumno=? and edc.ano=? and edc.sem=? AND ct.id_carga_total=? GROUP BY vp.nombre_completo", codPersona, ano, sem, idCargaTotal).Scan(&results2).Error

	return results2.Nombre, err

}

type ProfesorNombre struct {
	Nombre string `gorm:"column:nombre_completo" json:"nombre"`
}

// selecciona el plan actual de un alumno
// select_plan_actual(
func GetPlanAlumnoActualService(codPersona int) ([]models.PlanAlumnoActual, error) {
	db := config.Database
	var results []models.PlanAlumnoActual
	err := db.Raw("SELECT matricula.cod_tipo_matricula as cod_tipo_matricula, carrera_total.cod_especialidad,carrera_total.cod_mencion,carrera_total.cod_carrera,matricula.id_matricula, matricula.ano_matricula, matricula.sem_estado_final, matricula.ano_proceso, matricula.cod_estado_final, matricula.id_planes, carreras.nom_carrera, menciones.nom_mencion, matricula.ano_matricula, matricula.ano_ingreso, planes.anoplan, estado_final_matricula.descripcion, especialidades.nom_especialidad, planes.cod_tipo_plan, planes.cod_nivel_global, cu.nom_centro, carrera_total.cod_nivel_carrera FROM matricula INNER JOIN planes ON matricula.id_planes = planes.id_planes INNER JOIN carrera_total ON planes.id_carrera_total = carrera_total.id_carrera_total INNER JOIN centro_universitario cu ON carrera_total.cod_centro=cu.cod_centro INNER JOIN carreras ON carrera_total.cod_carrera = carreras.cod_carrera INNER JOIN menciones ON carrera_total.cod_mencion = menciones.cod_mencion INNER JOIN estado_final_matricula ON matricula.cod_estado_final = estado_final_matricula.cod_estado_final INNER JOIN especialidades ON carrera_total.cod_especialidad = especialidades.cod_especialidad WHERE (matricula.cod_persona = ?) ORDER BY matricula.ano_matricula DESC, matricula.id_matricula DESC", codPersona).Scan(&results).Error

	return results, err
}

// Selecciona la informacion de matricula de un alumno para un plan
// select_matricula_alumno(
func GetMatriculaAlumnoService(codPersona int, idPlan int) ([]models.MatriculaAlumno, error) {
	db := config.Database
	var results []models.MatriculaAlumno
	err := db.Raw("SELECT *,convert(varchar,fecha_matricula,110) as fecha_matricula1 FROM matricula INNER JOIN	estado_final_matricula ON matricula.cod_estado_final = estado_final_matricula.cod_estado_final WHERE (matricula.cod_persona = ?) AND (matricula.id_planes = ?) ORDER BY ano_matricula desc, matricula.cod_tipo_matricula desc;", codPersona, idPlan).Scan(&results).Error

	return results, err
}

// Selecciona la informacion del porcentaje de aprobacon de un alumno segun el plan certificado
// porcentaje_aprobacion_por_plan_certi(
func GetPorcentajeAprobacionPlanCertiService(codPersona int, idPlan int, anoIngreso int, anoProceso int) ([]models.PorcentajeAprobacionPlanCerti, error) {
	db := config.Database
	var results []models.PorcentajeAprobacionPlanCerti
	err := db.Raw("select  count(ua.id_malla) as ramos_aprobados , count(ma.id_malla) as ramos_malla, ((count(ua.id_malla)*100)/count(ma.id_malla)) as porce from   /* ramos rendidos nota mayor a 4*/ (SELECT     id_malla AS malla, id_malla = CASE WHEN (nota_final >= 4) THEN id_malla ELSE NULL END FROM notas_finales WHERE (cod_persona = ?) AND ano>=? AND ano<=?) ua right join /* ramos de la malla */ (select mp.id_malla,CASE WHEN mpp.id_malla IS NULL THEN mp.nro_ano ELSE mpp.nro_ano END AS nro_ano,CASE WHEN mpp.id_malla IS NULL THEN mp.nro_sem ELSE mpp.nro_sem END AS nro_sem from matricula m inner join malla_pionero mp on m.id_planes=mp.id_planes left join malla_pionero_personas mpp on mp.id_malla=mpp.id_malla and mpp.cod_persona=? where m.cod_persona=? and m.id_planes=? AND mp.cod_tipo_curso_malla<='2') ma on ua.id_malla=ma.id_malla group by ma.nro_sem,ma.nro_ano", codPersona, anoIngreso, anoProceso, codPersona, codPersona, idPlan).Scan(&results).Error
	return results, err
}

// Selecciona la informacion de la matricula del alumno, por el año de ingreso y proceso
// select_matricula_alumno_anoingreso_anoproceso(
func GetMatriculaAlumnoAnoIngresoAnoprocesoService(codPersona int, idPlan int, anoIngreso int, anoProceso int) ([]models.MatriculaAlumnoAnoIngresoAnoproceso, error) {
	db := config.Database
	var results []models.MatriculaAlumnoAnoIngresoAnoproceso
	err := db.Raw("SELECT *,convert(varchar,fecha_matricula,110) as fecha_matricula1, tipo_matricula.descripcion as tipo_matricula FROM matricula INNER JOIN estado_final_matricula ON matricula.cod_estado_final = estado_final_matricula.cod_estado_final INNER JOIN  tipo_matricula ON matricula.cod_tipo_matricula = tipo_matricula.cod_tipo_matricula WHERE (matricula.cod_persona = ?) AND (matricula.id_planes = ?) AND ano_ingreso=? AND ano_matricula=? ORDER BY ano_matricula desc, matricula.cod_tipo_matricula desc", codPersona, idPlan, anoIngreso, anoProceso).Scan(&results).Error
	return results, err
}

// Obtiene el nivel a partir del porcentaje de aprobacion.
// nivel_sem_cert(
func GetNivelSemCertService(porce []models.PorcentajeAprobacionPlanCerti, idPlan int) int {
	cont := 1
	nivel := 0

	for _, por := range porce {
		if por.Porce >= 80 {
			if idPlan > 0 {
				if (idPlan == 158 && cont == 5) || (idPlan == 341 && cont == 2) {
					cont++
				}
			}
			nivel = cont
		} else {
			break
		}
		cont++
	}

	if nivel != len(porce) {
		nivel++
	}

	return nivel
}

// Obtiene el semestre y año a partir del nivel.
// sem_ano(
func GetSemAnoService(nivel int) []string {
	carreras := []CarreraAnoSem{
		{Ano: "0", Sem: "0"},
		{Ano: "1", Sem: "1"},
		{Ano: "1", Sem: "2"},
		{Ano: "2", Sem: "1"},
		{Ano: "2", Sem: "2"},
		{Ano: "3", Sem: "1"},
		{Ano: "3", Sem: "2"},
		{Ano: "4", Sem: "1"},
		{Ano: "4", Sem: "2"},
		{Ano: "5", Sem: "1"},
		{Ano: "5", Sem: "2"},
		{Ano: "6", Sem: "1"},
		{Ano: "6", Sem: "2"},
		{Ano: "7", Sem: "1"},
		{Ano: "7", Sem: "2"},
	}

	arreglo := make([]string, 2)
	arreglo[0] = carreras[nivel].Ano
	arreglo[1] = carreras[nivel].Sem

	return arreglo
}

type CarreraAnoSem struct {
	Ano string
	Sem string
}

// Obtiene la cantidad de semestres del plan
// cantidad_semestres_plan(
func GetCantidadSemestresPlanService(idPlan int) (int, error) {
	db := config.Database
	var results []CantidadSemestres
	err := db.Raw("select count(*) as cantidad_semestres from (select nro_ano, nro_sem from sayd.malla_pionero where id_planes=? and cod_tipo_curso_malla<=2 group by nro_ano, nro_sem) t1", idPlan).Scan(&results).Error
	if len(results) > 0 {
		cant := results[0].Cantidad
		if idPlan == 158 {
			cant++
		}
		return cant, err
	} else {
		return 0, err
	}

}

type CantidadSemestres struct {
	Cantidad int `gorm:"column:cantidad_semestres" json:"cantidad"`
}

// Obtiene los parametros generales que indican la encuesta docente si esta habilitada
// encuesta_docente_habilitada(
func GetEncuestaDocenteHabilitadaService() (bool, error) {
	db := config.Database
	var results []models.ParametrosGenerales
	err := db.Raw("select * from parametros_generales where cod_tipo_parametro=30 and switch=1").Scan(&results).Error
	if len(results) > 0 {
		return true, err
	}
	return false, err
}

// obtiene las encuestas docentes por contestar
// encuestas_por_contestar(
func GetEncuestaPorContestarService(codPersona string, ano string, sem string) ([]models.EncuestaDocentePorContestar, error) {
	db := config.Database
	var results []models.EncuestaDocentePorContestar
	var cond11 string
	var cond12 string
	if sem == "2" {
		cond11 = "/*  anuales */ UNION SELECT ctp.id_carga_total from enc_docente_contestadas edc INNER JOIN carga_total_profesores ctp ON edc.id_carga_total_profesor=ctp.id_carga_total_profesor INNER JOIN usuarios_aula ua ON ua.id_carga_total=ctp.id_carga_total AND ua.cod_persona=" + codPersona + " INNER JOIN malla_pionero mp ON ua.id_malla=mp.id_malla WHERE edc.cod_alumno=" + codPersona + " AND edc.ano=" + ano + " and mp.duracion='A'"
		cond12 = "UNION /*  anuales */ SELECT ua.id_carga_total, ua.id_usuario_aula FROM usuarios_aula ua INNER JOIN malla_pionero mp on ua.id_malla=mp.id_malla INNER JOIN asistencia_alumnos aa ON ua.id_usuario_aula=aa.id_usuario_aula INNER JOIN carga_total_profesores ctp ON ua.id_carga_total=ctp.id_carga_total WHERE (ua.ano = " + codPersona + ") AND (ua.sem = '1') AND mp.duracion='A' AND ua.cod_persona=" + codPersona + " AND (ua.cod_estado_inscripcion=1 OR ua.cod_estado_inscripcion=2 OR ua.cod_estado_inscripcion=7 OR ua.cod_estado_inscripcion=12) AND aa.porcentaje>=60 AND ctp.cod_persona!='30283'	GROUP BY ua.id_carga_total, ua.id_usuario_aula"
	} else {
		cond11 = ""
		cond12 = ""
	}
	err := db.Raw("/*  encuesta contestadas    */SELECT * FROM (SELECT id_carga_total from enc_docente_contestadas edc INNER JOIN carga_total_profesores ctp ON edc.id_carga_total_profesor=ctp.id_carga_total_profesor WHERE edc.cod_alumno=? AND ano=? AND sem=? "+cond11+" ) t1 RIGHT JOIN /*  encuesta que debe contestar */ (SELECT ua.id_carga_total, ua.id_usuario_aula FROM usuarios_aula ua INNER JOIN asistencia_alumnos aa ON ua.id_usuario_aula=aa.id_usuario_aula INNER JOIN carga_total_profesores ctp ON ua.id_carga_total=ctp.id_carga_total WHERE (ua.ano = ?) AND (ua.sem = ?) AND (ua.cod_persona = ?) AND (ua.cod_estado_inscripcion=1 OR ua.cod_estado_inscripcion=2 OR ua.cod_estado_inscripcion=7 OR ua.cod_estado_inscripcion=12) AND aa.porcentaje>=60 AND ctp.cod_persona!='30283' GROUP BY ua.id_carga_total, ua.id_usuario_aula "+cond12+") t2 ON t1.id_carga_total=t2.id_carga_total WHERE (t1.id_carga_total IS NULL)", codPersona, ano, sem, ano, sem, codPersona).Scan(&results).Error
	return results, err

}
