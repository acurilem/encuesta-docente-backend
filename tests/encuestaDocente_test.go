package aceptacion

import (
	"fmt"
	"testing"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type MyData struct {
	Token string `json:"token"`
}

func TestMain(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic("Error loading .env.test file")
	}
}

func getToken(login string, contrasena string) (string, error) {
	// Define the URL to which you want to send the HTTP POST request
	url := os.Getenv("ADDR") + "/auth/login"

	// Define the request body as a JSON string
	requestBody := []byte(`{
			"username": "` + login + `",
			"password": "` + contrasena + `"
		}`)

	responseBody, shouldReturn := callApi(url, "", "POST", bytes.NewBuffer(requestBody))
	if shouldReturn {
		return "", nil
	}

	// Convert the response body to a string and print it
	var data MyData

	// Handle the response as needed
	err := json.Unmarshal([]byte(string(responseBody)), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "", err
	}

	return data.Token, nil
}

func callApi(url string, token string, method string, requestBody io.Reader) (string, bool) {
	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request with a Bearer token in the Authorization header
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", true
	}

	// Set the content type header to indicate JSON
	req.Header.Set("Content-Type", "application/json")

	// Add the Bearer token to the request header
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", true
	}
	defer response.Body.Close()

	// Read and display the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", true
	}

	// Convert the response body to a string and print it
	responseBody := string(body)
	return responseBody, false
}

func TestIndexAlumno(t *testing.T) {

	token, err := getToken("21645053-1", "secret")
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}

	url := os.Getenv("ADDR") + "/api/v1/encuestaDocenteRespuestas/indexAlumno"

	responseBody, shouldReturn := callApi(url, token, "GET", nil)
	if shouldReturn {
		return
	}

	// Verificar información esperada
	substrings := []string{
		"\"idMalla\":26899",
		"\"nomGrupo\":\"COMUNICACIÓN EFECTIVA.\"",
		"\"codTipoSeccion\":1",
		"\"idMalla\":26894",
		"\"nomGrupo\":\"ETICA Y JUSTICIA\"",
		"\"codTipoSeccion\":1",
		"\"idMalla\":26896",
		"\"nomGrupo\":\"HISTORIA DEL DERECHO\"",
		"\"codTipoSeccion\":1",
		"\"idMalla\":26897",
		"\"nomGrupo\":\"NIVEL DE COMPETENCIAS I\"",
		"\"codTipoSeccion\":1",
		"\"idMalla\":26893",
		"\"nomGrupo\":\"TEORIA DEL DERECHO\"",
		"\"codTipoSeccion\":1",
		"\"idMalla\":26895",
		"\"nomGrupo\":\"TEORÍA POLÍTICA DEL ESTADO\"",
		"\"codTipoSeccion\":1",
		"\"IdCargaTotal\":56323",
		"\"Porcentaje\":100",
		"\"IdCargaTotal\":56012",
		"\"Porcentaje\":97",
		"\"IdCargaTotal\":56013",
		"\"Porcentaje\":90",
		"\"IdCargaTotal\":56015",
		"\"Porcentaje\":100",
		"\"IdCargaTotal\":56017",
		"\"Porcentaje\":0",
	}

	for _, substring := range substrings {
		assert.Contains(t, responseBody, substring)
	}

}

func TestEncuestaDocente(t *testing.T) {

	token, err := getToken("15308450-5", "secret")
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}

	url := os.Getenv("ADDR") + "/api/v1/encuestaDocenteRespuestas/"

	formData := gin.H{
		"idCargaTotalProfesor": 12,
		"p1":                   1,
		"p2":                   2,
		"p3":                   3,
		"p4":                   1,
		"p5":                   2,
		"p6":                   3,
		"p7":                   1,
		"p8":                   2,
		"nota":                 6,
		"comentario":           "comentario 1",
		"ano":                  1998,
		"sem":                  1,
	}

	// Convierte el formulario a un reader
	jsonData, err := json.Marshal(formData)
	if err != nil {
		// Manejar el error de serialización
		return
	}

	encuesta := bytes.NewReader(jsonData)

	responseBody, shouldReturn := callApi(url, token, "POST", encuesta)
	if shouldReturn {
		return
	}

	substring := "\"idCargaTotalProfesor\":12"
	assert.Contains(t, responseBody, substring)
	substring = "\"p1\":1"
	assert.Contains(t, responseBody, substring)
	substring = "\"p2\":2"
	assert.Contains(t, responseBody, substring)
	substring = "\"p3\":3"
	assert.Contains(t, responseBody, substring)
	substring = "\"p4\":1"
	assert.Contains(t, responseBody, substring)
	substring = "\"p5\":2"
	assert.Contains(t, responseBody, substring)
	substring = "\"p6\":3"
	assert.Contains(t, responseBody, substring)
	substring = "\"p7\":1"
	assert.Contains(t, responseBody, substring)
	substring = "\"p8\":2"
	assert.Contains(t, responseBody, substring)
	substring = "\"nota\":6"
	assert.Contains(t, responseBody, substring)
	substring = "\"comentario\":\"comentario 1\""
	assert.Contains(t, responseBody, substring)
	substring = "\"ano\":1998"
	assert.Contains(t, responseBody, substring)
	substring = "\"sem\":1"
	assert.Contains(t, responseBody, substring)

}
