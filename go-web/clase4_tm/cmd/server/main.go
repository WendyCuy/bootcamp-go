package main

/* Importación de dependencias. El main se encontrará en el directorio cmd/server.
Importamos las dependencias necesarias. */

import (
	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tt/cmd/server/handler"
	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tt/internal/products"
	"github.com/WendyCuy/bootcamp-go/go-web/clase3_tt/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/* Main del programa.  Instanciamos cada capa del dominio Productos y utilizaremos
los métodos del controlador para cada endpoint.

Enviar Base de datos al repositorio.  Instanciamos desde el Factory de store
indicando el tipo archivo (FileType) y donde deseamos guardar el json,
y le pasamos la base de datos al repositorio.*/

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "products.json")
	repo := products.NewRepository(db)
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
