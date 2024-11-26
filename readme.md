## estrutura do projeto

go-backend/
├── main.go
├── config/
│   └── config.go    
├── controllers/
│   └── cliente_controller.go    
├── docs/        
├── models/
│   └── cliente.go     
├── routes/
│   └── cliente.go  
│   └── routes.go           
├── tests/
│   └── cliente_test.go      
└── go.mod 

## criando o projeto
go mod init go-backend

## dependencias necessárias
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/gorilla/mux

## incluindo o swagger
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

## adicionando o swagger
swag init

## instalação de teste automatizado
go get github.com/stretchr/testify
go get github.com/gin-gonic/gin

## executando os testes
go test ./tests

## Para executar 
go run main.go