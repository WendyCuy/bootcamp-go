package main

/* Importación de dependencias. El main se encontrará en el directorio cmd/server.
Importamos las dependencias necesarias. */

import (
	"fmt"
	"os"

	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/cmd/server/handler"
	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/docs"
	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/internal/products"
	"github.com/WendyCuy/bootcamp-go/go-testing/clase2_tt-test/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/* Main del programa.  Instanciamos cada capa del dominio Productos y utilizaremos
los métodos del controlador para cada endpoint.

Enviar Base de datos al repositorio.  Instanciamos desde el Factory de store
indicando el tipo archivo (FileType) y donde deseamos guardar el json,
y le pasamos la base de datos al repositorio.*/

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	//Documentacion de swagger.
	// Al ingresar a http://localhost:8080/docs/index.html se puede visualizar la documentación
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	// Se realiza ajuste con validador de golangci-lint
	err := r.Run()
	if err != nil {
		fmt.Println("No se pudo iniciar el programa")
	}

}

//go run cmd/server/main.go
//
