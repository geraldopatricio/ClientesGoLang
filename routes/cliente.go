package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Lista todos os clientes
// @Description Retorna uma lista de todos os clientes registrados
// @Tags Clientes
// @Accept json
// @Produce json
// @Success 200 {array} models.Cliente
// @Router /clientes [get]
func GetAllClientes(c *gin.Context) {
	controllers.GetAllClientes(c)
}

// @Summary Cria um novo cliente
// @Description Adiciona um novo cliente no sistema
// @Tags Clientes
// @Accept json
// @Produce json
// @Param cliente body models.Cliente true "Dados do Cliente"
// @Success 201 {object} models.Cliente
// @Router /clientes [post]
func CreateCliente(c *gin.Context) {
	controllers.CreateCliente(c)
}

// @Summary Busca cliente por ID
// @Description Retorna os dados de um cliente pelo ID
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} models.Cliente
// @Failure 404 {string} string "Cliente não encontrado"
// @Router /clientes/{id} [get]
func GetClienteByID(c *gin.Context) {
	controllers.GetClienteByID(c)
}

// @Summary Atualiza um cliente
// @Description Atualiza os dados de um cliente pelo ID
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Param cliente body models.Cliente true "Dados atualizados do Cliente"
// @Success 200 {object} models.Cliente
// @Failure 404 {string} string "Cliente não encontrado"
// @Router /clientes/{id} [put]
func UpdateCliente(c *gin.Context) {
	controllers.UpdateCliente(c)
}

// @Summary Exclui um cliente
// @Description Remove um cliente pelo ID
// @Tags Clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 204 {string} string "Cliente removido com sucesso"
// @Failure 404 {string} string "Cliente não encontrado"
// @Router /clientes/{id} [delete]
func DeleteCliente(c *gin.Context) {
	controllers.DeleteCliente(c)
}
