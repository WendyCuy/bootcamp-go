package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id        string
	Nombre    string
	Color     string
	Precio    int
	Stock     int
	Codigo    string
	Publicado bool
	Fecha     string
}

func main() {
	router := gin.Default()

	// Captura la solicitud GET "Hola nombre"
	router.GET("hola/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hola " + name,
		})
	})

	datosJson, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	productos := Productos{}
	err = json.Unmarshal(datosJson, &productos)
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/productos", func(ctx *gin.Context) {
		ctx.JSON(200, productos)
	})

	router.Run(":8080")

}
