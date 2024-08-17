# UMAG Template Backend Golang

## Guía inicio

### Cambiar nombre paquete

Lo primero a realizar es cambiar el nombre del paquete de go para comenzar a trabajar. El nombre actual del paquete es `github.com/citiaps/SID-UMAG-Encuesta-docente-backend`. Lo que se debe hacer entonces, es cabiar el nombre al nombre del repositorio en github en donde va a residir el proyecto.

### Variables de entorno

Cree un archivo `.env` en la raiz del proyecto:

```txt
- config
- controller
- ...
- .env
```

Copie el contenido de `.env.example` dentro del archivo recientemente creado y reemplace las variables de entorno según convenga.

### Correr proyecto con Docker

Dentro de los archivos del template, existe un Dockerfile llamado `Dockerfile.dev`, este construye el entorno de desarrollo para trabajar, además de venir instalado un demonio para automáticamente
hacer re build y re run del proyecto en caso de existir un cambio
en los archivos del mismo.

### Swagger

Es imperante el documentar, por esto mismo, a continuación se muestra
una introducción y guía de uso de la herramienta Swagger. Herramienta
con el fin de documentar APIs a través de comentarios de código.

Documentación oficial del formato de los comentarios declarativos para documentar: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

La versión de Swagger de la que se va a hacer uso es una implementación para el framework Gin-Gonic de Go, la cual es llamada
[`gin-swagger`](https://github.com/swaggo/gin-swagger).

#### Instalación

La herramienta, a nivel de paquete, ya viene instalada dentro de los paquetes en el proyecto, adicionalmente, debe instalar la herramienta `Swag` a nivel de sistema con el siguiente comando:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Compruebe la instalación con el comando:

```
swag
```

Si se encuentra con el error "command not found", configure bien su `GOPATH` dentro de su sistema. En caso de usar linux, añade estas dos lineas en su `~/.bashrc` o `~/.zhsrc`

```bash
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

Y luego corra:

```bash
source ~/.bashrc (o ~/.zhsrc)
```

#### Inicialización

A continuación corra el comando:

```bash
swag init
```

Si no hubo ningún error, perfecto, está instalado e inicializado swagger en el proyecto.

#### Configuración en código

Se necesitan hacer tres imports en el código donde se ejecuta la función `main`, que es el mismo lugar donde se ejecuta el router de gin-gonic, los tres imports son:

```go
import (
	"github.com/citiaps/SID-UMAG-Encuesta-docente-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
```

Nota: En este caso se usa `github.com/citiaps/SID-UMAG-Encuesta-docente-backend` como nombre de paquete, pero esto varia según repositorio/proyecto.

Se importa entonces en primera instancia el paquete `docs` que creó swag al correr `swag init`. Luego, `swaggerFiles` swagger embed files. Y por último el `ginSwagger`, paquete que wrapea el handler del paquete original de swagger, `swaggo`.

Estos imports ya están implementados en código de la siguiente forma:

```go
// Docs
docs.SwaggerInfo.BasePath = "/api/v1/x"
docs.SwaggerInfo.Version = "v1"
docs.SwaggerInfo.Host = "localhost:8000"
// Route docs
app.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

En el bloque *Docs* se setea la información básica mínima para comenzar a utilizar el paquete de Swagger.

Al iniciar un nuevo proyecto a partir de este template, necesita añadir y modificar esta información, como la versión última de la API, la `basepath`, entre otros parametros, hagalo según documentación y usando estas tres configuración mínimas requeridas.

Para el bloque *Route docs*, se setea la ruta de tipo `GET` para acceder a los documentos de Swagger, tanto a la web como a los archivos JSON. 

```go
// [GET] /api/v1/x/swagger/index.html <- Accede a la interfaz gráfica web
// [GET] /api/v1/x/swagger/docs.json <- Accede a los documentos JSON autogenerados
```

#### Documentar en Swagger

Como se habló anteriormente, Swaggo es una herramienta declarativa que permite documentar una API con comentarios.

La declaración de estos comentarios deben ser dados en dos lados, primeramente, en la función `main`, se declaran comentarios sobre información de su API como por ejemplo, qué datos produce, qué etiquetas existen, el titulo, versión, entre otros datos. En este template se encuentra una implementación completa de todo lo que debe declarar de información para su API, modifiquela según proyecto.

A continuación se hace expliración de cada parametro configurado en este template:

```go
// @title          Nombre de la API
// @version        1.0
// @description    Descripción de la API
// @termsOfService http://swagger.io/terms/

// @contact.name  Nombre de contacto organización
// @contact.url   URL de contacto organización
// @contact.email Email de contacto organización

// lincense.name  Licencia del programa
// @license.url URL de la licencia para el programa

// @tag.name        Nombre etiqueta principal
// @tag.description Descripción etiqueta principal

// @host     Host
// @BasePath Path base

/*
    Todo el bloque de abajo habla sobre las definiciones de seguridad. En este caso
    se refiere a una autenticación basada en Bearer Token JWT por medio del header
    Authorization.
*/
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                BearerJWTToken in Authorization Header



/*
    Puede agregar tantos "@accept" o "@produce" como tipos de MIME que las API pueden producir o aceptar.
    Ej:
    @accept json
    @accept xml
*/
// @product octet-stream
// @product application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @product application/pdf
// @accept  Tipo de MIME que acepta la API
// @produce Tipo de MIME que produce la API

// @schemes http https
```

El segundo lugar donde usted debe documentar, en el las funciones de controlador que manejan las peticiones. Ejemplo de un controlado y cómo documentar:

```go
package controller

import (
	"github.com/gin-gonic/gin"
)

var db = map[string]string{
    "1": "Pelusa",
    "2": "Princesa",
}

// GetDog godoc
// @Summary     Obtener un perro por su ID
// @Description Esta funcionalidad toma un idDog como parametro de ruta, para luego tratar de encontrarlo en la base de datos, si falla arroja un NotFound y si tiene exito retorna el perro.
// @Param       idDog path string true "ID Del perro a consultar" 
// @Tags        dogs
// @Accept      json
// @Product     json
// @Success     200 {object} models.Dog
// @Failure     404 {object} res.Response{} "Dog not found"
// @Router      /api/v1/dogs/{idDog} [get]
func GetDog(ctx *gin.Context) {
    idDog := ctx.GetParam("idDog")
    if dog, ok := db[idDog]; !ok {
        ctx.AbortWithStatusJSON(
            http.StatusNotFound,
            gin.H{
                "message": "Dog not found",
            },
        )
        return
    }

	ctx.JSON(http.StatusOK, dog)
}
```

Como puede ver, justo encima de la funcionalidad que maneja una petición de tipo `GET`, la cual podemos imaginar responde a la ruta `/api/v1/dogs/:idDog`, está documentando la API por medio de comentarios, veamos más de cerca estos comentarios para entenderlos uno a uno.

```go
// GetDog godoc <- Nombre función más la palabra godoc para identificar que es documentación de Swagger para la API
// @Summary     Obtener un perro por su ID <- Resumen de la funcionalidad
// @Description Esta funcionalidad toma un idDog como parametro de ruta, para luego tratar de encontrarlo en la base de datos, si falla arroja un NotFound y si tiene exito retorna el perro. <- Explicación a detalle de la funcionalidad
// @Param       idDog path string true "ID Del perro a consultar" <- Parametros de consulta, datos requeridos, etc. Generalmente se usa para requerir el body, path param, query param, entre otros
// @Tags        dogs <- Etiqueta o etiquetas del endpoint
// @Accept      json <- MIME que acepta
// @Produce     json <- MIME que produce
// @Success     200 {object} models.Dog <- Qué estado, tipo de dato y módelo de dato se genera en caso de exito
// @Failure     404 {object} res.Response{} "Dog not found" <- Error o errores que produce la API
// @Router      /api/v1/dogs/{idDog} [get] <- Path del endpoint y el verbo HTTP
```

Como mínimo, la documentación debe tener:

- Nombre de función más godoc: `{{func}} godoc`
- Resumen de la funcionalidad: `@Summary`
- Descripción de la funcionalidad: `@Description`
- Etiqueta: `@Tags` 
- MIME(s) que acepta: `@Accept`
- MIME(s) que produce: `@Produce`
- Estado de exito: `@Success`
- Posibles respuestas de error: `@Failure`
- Path y verbo: `@Router`

Otras *API Operation* se deben documentar según caso de uso. La documentación de estas, se encuentra al inicio del bloque [Swagger](#swagger).

#### Actualizar Swagger

Cuando quiera actualizar los documentos de Swagger, ejecute los siguientes comandos:

```bash
swag fmt # Formatea los swag comments
swag init # Genera los documentos de Swagger nuevamente
```
