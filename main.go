package main

import (
	"go-backend/config"
	"go-backend/routes"

	_ "go-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Clientes GoLang
// @version 1.0
// @description API para gerenciar clientes (CRUD)
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte
// @contact.url http://gpsoft.com.br
// @contact.email geraldo@gpsoft.com.br

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
