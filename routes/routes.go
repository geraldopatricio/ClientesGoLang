package routes

import (
	"go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	clienteGroup := r.Group("/clientes")
	{
		clienteGroup.GET("/", controllers.GetAllClientes)
		clienteGroup.POST("/", controllers.CreateCliente)
		clienteGroup.GET("/:id", controllers.GetClienteByID)
		clienteGroup.PUT("/:id", controllers.UpdateCliente)
		clienteGroup.DELETE("/:id", controllers.DeleteCliente)
	}
}
