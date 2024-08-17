package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/models"
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/services"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateEncuestaDocenteRespuestas godoc
//
//	@Summary		Guarda las respuestas de una encuesta de un alumno hacia un profesor.
//	@Description	Esta funcionalidad recibe un body con los datos del profesor y los puntajes de la encuesta asignados por el alumno, son guardados en la BD y se deja registro de que la encuesta ya fue contestada, además se deja rastro de estos Create.
//	@Param			data	body	models.EncuestaDocenteRespuestas	true	"Datos de la encuesta respondida por el alumno"
//	@Tags			EncuestaDocenteRespuestas
//	@Accept			json
//	@Product		json
//	@Success		201	{object}	models.EncuestaDocenteRespuestas
//	@Failure		400	{object}	ErrorResponse	"Petición erronea"
//	@Security		Bearer
//	@Router			/encuestaDocenteRespuestas [post]
func CreateEncuestaDocenteRespuestas(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion}
	usuario, err := services.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetUser:\n" + err.Error()})
		return
	}
	codPersona := int(usuario.ID)
	var encuestaDocenteRespuestas models.EncuestaDocenteRespuestas
	//En caso de algun error
	if err := ctx.ShouldBindJSON(&encuestaDocenteRespuestas); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el EncuestaDocenteRespuestas
	createdEncuestaDocenteRespuestas, err := services.CreateEncuestaDocenteRespuestasService(codPersona, encuestaDocenteRespuestas)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el EncuestaDocenteRespuestas"})
		return
	}

	// Devuelve el mensaje de confirmación y las respuestas guardardas
	log.Println("EncuestaDocenteRespuestas creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdEncuestaDocenteRespuestas)
}

// IndexAlumno godoc
//
//	@Summary		Obtiene información del alumno, las asignaturas cursadas en el actual semestre y año, y los profesores de cada asignatura.
//	@Description	Esta funcionalidad obtiene de la base de datos el año y semestre actual, luego con el codPersona se obtienen los datos del alumno y sus asignaturas, luego con el idCargaTotal, se obtienen los profesores y si las encuestas ya fueron contestadas.
//	@Tags			IndexAlumno
//	@Accept			json
//	@Product		json
//	@Success		200	{object}	models.IndexAlumno
//	@Failure		400	{object}	ErrorResponse	"Petición erronea"
//	@Security		Bearer
//	@Router			/encuestaDocenteRespuestas/indexAlumno [get]
//
// index_alumno()
func GetIndexAlumno(ctx *gin.Context) {
	usuario, err := services.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetUser"})
		return
	}
	codPersona := int(usuario.ID)

	var retorno models.IndexAlumno
	anoProceso, err := services.GetAllAnoProcesoService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetAllAnoProcesoService"})
		return
	}

	ano, err := strconv.Atoi(strings.TrimSpace(anoProceso[0].AnoProc))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir los datos del GetAllAnoProcesoService"})
		return
	}
	semestre := "1"
	if anoProceso[0].SemProc == "1" {
		ano -= 1
		semestre = "2"
	}
	anoAux := strconv.Itoa(ano)

	asignaturasInscritas, err := services.GetSeccionesInscritasAlumnoService(anoAux, semestre, codPersona)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetSeccionesInscritasAlumnoService"})
		return
	}
	retorno.AsignaturasInscritas = asignaturasInscritas
	var asistenciaAsignaturas []models.AsistenciaAsign
	var contestadas []models.Contestada
	var profesorSeleccionado []models.ProfesorSeleccionado
	var profesores []models.ProfesoresResumen
	if len(asignaturasInscritas) > 0 {
		i := 0
		for _, asigIns := range asignaturasInscritas {
			var asistenciaAux models.AsistenciaAsign
			var contestadaAux models.Contestada
			var profesorSelecAux models.ProfesorSeleccionado
			var profesoresAux models.ProfesoresResumen
			idCargaTotal := asigIns.IdCargaTotal
			idUsuarioAula := asigIns.IdUsuarioAula
			asistenciaAlumno, err := services.GetValidarAsistenciaService(idUsuarioAula)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetValidarAsistenciaService"})
				return
			}
			if len(asistenciaAlumno) > 0 {
				asistenciaAux.IdCargaTotal = idCargaTotal
				asistenciaAux.Porcentaje = asistenciaAlumno[0].Porcentaje
			} else {
				asistenciaAux.IdCargaTotal = idCargaTotal
				asistenciaAux.Porcentaje = 0
			}

			// Profesores de un idCargaTotal
			profesoresAux.IdCargaTotal = idCargaTotal
			profesoresAux.CargaTotalProfesor, err = services.GetProfesoresCargaTotalService(idCargaTotal)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetProfesoresCargaTotalService"})
				return
			}
			contestadaAux.IdCargaTotal = idCargaTotal
			contestadaAux.Flag, err = services.GetValidarEncuestasContestadasService(idCargaTotal, ano, semestre, codPersona)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetValidarEncuestasContestadasService"})
				return
			}
			profesorSelecAux.IdCargaTotal = idCargaTotal
			profesorSelecAux.Nombre, err = services.GetProfesorSeleccionadoService(idCargaTotal, ano, semestre, codPersona)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetProfesorSeleccionadoService"})
				return
			}
			i = i + 1
			asistenciaAsignaturas = append(asistenciaAsignaturas, asistenciaAux)
			profesores = append(profesores, profesoresAux)
			contestadas = append(contestadas, contestadaAux)
			profesorSeleccionado = append(profesorSeleccionado, profesorSelecAux)
		}
		retorno.AsistenciaAsignaturas = asistenciaAsignaturas
		retorno.Profesores = profesores
		retorno.Contestadas = contestadas
		retorno.ProfesorSeleccionado = profesorSeleccionado

	}
	planAlumno, err := services.GetPlanAlumnoAnoService(codPersona, ano)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetPlanAlumnoAnoService"})
		return
	}
	retorno.PlanAlumno = planAlumno

	nivelAlumno := 1

	planAlumnoActual, err := services.GetPlanAlumnoActualService(codPersona)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetPlanAlumnoActualService"})
		return
	}
	idPlan := planAlumnoActual[0].IdPlanes
	retorno.InfoPlan = planAlumnoActual
	matriculaAlumnoActual, err := services.GetMatriculaAlumnoService(codPersona, idPlan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetMatriculaAlumnoService"})
		return
	}
	if len(matriculaAlumnoActual) > 0 {
		anoProcesoNum := matriculaAlumnoActual[0].AnoIngreso
		anoIngreso := matriculaAlumnoActual[0].AnoMatricula
		retorno.AnoProcesoNum = anoProcesoNum
		retorno.AnoIngreso = anoIngreso
		porce, err := services.GetPorcentajeAprobacionPlanCertiService(codPersona, idPlan, anoIngreso, anoProcesoNum)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetPorcentajeAprobacionPlanCertiService"})
			return
		}
		nivelSem := services.GetNivelSemCertService(porce, idPlan)
		ubicacion := services.GetSemAnoService(nivelSem)
		infoMatricula, err := services.GetMatriculaAlumnoAnoIngresoAnoprocesoService(codPersona, idPlan, anoIngreso, anoProcesoNum)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetMatriculaAlumnoAnoIngresoAnoprocesoService"})
			return
		}
		estadoFinalMatricula := 0
		if len(infoMatricula) > 0 {
			estadoFinalMatricula = infoMatricula[0].CodEstadoFinal
		}

		if estadoFinalMatricula == 4 || estadoFinalMatricula == 5 || estadoFinalMatricula == 6 {
			switch estadoFinalMatricula {
			case 4: // titulado
				nivelAlumno = 9
			case 5: // egresado
				nivelAlumno = 10
			case 6: // memorista
				nivelAlumno = 11
			}
		} else {
			cantidadSem, err := services.GetCantidadSemestresPlanService(idPlan)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del GetCantidadSemestresPlanService"})
				return
			}
			infoCantidad := services.GetSemAnoService(cantidadSem)
			cantidadSem, err = strconv.Atoi(infoCantidad[0])
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir infoCantidad"})
				return
			}
			ubi, err := strconv.Atoi(ubicacion[0])
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir ubicacion"})
				return
			}
			if ubi > cantidadSem {
				ubi = cantidadSem
			}
			if cantidadSem == 0 {
				nivelAlumno = 1
			} else {
				nivelAlumno = ubi
			}

		}

	}
	if nivelAlumno != 0 {
		retorno.NivelContesta = 0
	} else {
		retorno.NivelContesta = 1
	}
	// Devuelve el mensaje de confirmación y los datos del alumno
	ctx.JSON(http.StatusOK, retorno)
}

// GetConfirmacionEncuestasPorContestarRespuestas godoc
//
//	@Summary		Obtiene información del alumno, el actual semestre y año actual, y retorna un bool indicando si tiene encuestas docentes por contestar.
//	@Description	Esta funcionalidad obtiene de la base de datos el año y semestre actual, luego con el codPersona y el semestre directamente anterior se obtienen las encuestas docentes que faltan por contestar, si se obtiene 1 o más, se retorna un true, en caso contrario un false.
//	@Tags			ConfirmacionEncuestasPorContestarRespuestas
//	@Accept			json
//	@Product		json
//	@Success		200	{object}	PorContestar
//	@Failure		400	{object}	ErrorResponse	"Petición erronea"
//	@Security		Bearer
//	@Router			/encuestaDocenteRespuestas/porContestar [get]
func GetConfirmacionEncuestasPorContestarRespuestas(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	token := services.GetToken(ctx)
	user, err := services.ValidateJWTToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error en validar token en controlador GetVerHorarioAlumnoPDF": err.Error()})
		return
	}
	var codPersona = uint(user["cod_persona"].(float64))

	// Si el usuario es calara entonces retorna true
	if codPersona == 43284 {
		var r PorContestar
		r.Response = true
		ctx.JSON(http.StatusCreated, gin.H{"response": r.Response})
		return
	}

	// Se llama al servicio que crea el EncuestaDocenteRespuestas
	respuestasAnoProceso, err := services.GetAllAnoProcesoService()
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener GetAllAnoProcesoService"})
		return
	}
	var sem string
	var ano string
	if respuestasAnoProceso[0].SemProc == "2" {
		sem = "1"
		ano = respuestasAnoProceso[0].AnoProc
	} else {
		sem = "2"
		anoAux, _ := strconv.Atoi(respuestasAnoProceso[0].AnoProc)
		ano = strconv.Itoa(anoAux - 1)
	}
	// Se llama al servicio EncuestaDocenteRespuestas
	respuestasEncuestaPorContestar, err := services.GetEncuestaPorContestarService(strconv.Itoa(int(codPersona)), ano, sem)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener GetEncuestaPorContestarService"})
		return
	}
	// Se llama al servicio EncuestaDocenteHabilitada
	respuestasEncuestaDocenteHabilitada, err := services.GetEncuestaDocenteHabilitadaService()
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener GetEncuestaDocenteHabilitadaService"})
		return
	}
	retorno := false
	if len(respuestasEncuestaPorContestar) > 0 {
		if respuestasEncuestaDocenteHabilitada {
			retorno = true
		}
	}
	// Devuelve el mensaje de confirmación
	var r PorContestar
	r.Response = retorno
	ctx.JSON(http.StatusCreated, gin.H{"response": r.Response})
}

type PorContestar struct {
	Response bool `json:"response"`
}
