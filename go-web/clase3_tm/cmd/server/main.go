package main

/* Importación de dependencias. El main se encontrará en el directorio cmd/server.
Importamos las dependencias necesarias. */

import (
	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tm/cmd/server/handler"
	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tm/internal/products"
	"github.com/gin-gonic/gin"
)

/* Main del programa.  Instanciamos cada capa del dominio Productos y utilizaremos
los métodos del controlador para cada endpoint. */

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}

//go run cmd/server/main.go
