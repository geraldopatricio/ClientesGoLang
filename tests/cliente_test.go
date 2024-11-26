package tests

import (
	"bytes"
	"encoding/json"
	"go-backend/controllers"
	"go-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/clientes", controllers.GetAllClientes)
	r.POST("/clientes", controllers.CreateCliente)
	r.GET("/clientes/:id", controllers.GetClienteByID)
	r.PUT("/clientes/:id", controllers.UpdateCliente)
	r.DELETE("/clientes/:id", controllers.DeleteCliente)
	return r
}

func TestCreateCliente(t *testing.T) {
	cliente := models.Cliente{
		Nome:     "João da Silva",
		CPF:      "12345678900",
		Zap:      "11987654321",
		Tel:      "1132345678",
		Endereco: "Rua Teste, 123",
		Numero:   "123",
		Bairro:   "Centro",
		Cidade:   "São Paulo",
		UF:       "SP",
	}

	jsonData, _ := json.Marshal(cliente)

	w := performRequest("POST", "/clientes", jsonData)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "João da Silva")
}

func TestGetAllClientes(t *testing.T) {
	w := performRequest("GET", "/clientes", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String())
}

func TestGetClienteByID(t *testing.T) {
	cliente := models.Cliente{
		Nome:     "Maria Oliveira",
		CPF:      "09876543211",
		Zap:      "11987654321",
		Tel:      "1132345678",
		Endereco: "Avenida Teste, 456",
		Numero:   "456",
		Bairro:   "Centro",
		Cidade:   "São Paulo",
		UF:       "SP",
	}

	jsonData, _ := json.Marshal(cliente)
	performRequest("POST", "/clientes", jsonData)

	w := performRequest("GET", "/clientes/1", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Maria Oliveira")
}

func TestUpdateCliente(t *testing.T) {
	cliente := models.Cliente{
		Nome:     "Carlos Souza",
		CPF:      "11223344556",
		Zap:      "11976543210",
		Tel:      "1135678901",
		Endereco: "Rua Atualizada, 789",
		Numero:   "789",
		Bairro:   "Vila Teste",
		Cidade:   "São Paulo",
		UF:       "SP",
	}

	jsonData, _ := json.Marshal(cliente)

	w := performRequest("PUT", "/clientes/1", jsonData)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Carlos Souza")
}

func TestDeleteCliente(t *testing.T) {
	cliente := models.Cliente{
		Nome:     "Luiza Santos",
		CPF:      "99887766544",
		Zap:      "11923456789",
		Tel:      "1145678901",
		Endereco: "Rua Teste deleta, 321",
		Numero:   "321",
		Bairro:   "Centro",
		Cidade:   "São Paulo",
		UF:       "SP",
	}

	jsonData, _ := json.Marshal(cliente)
	performRequest("POST", "/clientes", jsonData)

	w := performRequest("DELETE", "/clientes/2", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Cliente deletado com sucesso")
}

func performRequest(method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(w, req)
	return w
}
