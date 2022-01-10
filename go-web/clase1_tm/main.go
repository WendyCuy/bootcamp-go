/* Ejercicio 1 - Estructura un JSON
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
1. Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
2. Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

Ejercicio 2 - Hola {nombre}
1. Crea dentro de la carpeta go-web un archivo llamado main.go
2. Crea un servidor web con Gin que te responda un JSON que tenga una clave
“message” y diga Hola seguido por tu nombre.
3. Pegale al endpoint para corroborar que la respuesta sea la correcta.

Ejercicio 3 - Listar Entidad
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
1. Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
2. Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
3. Genera un handler para el endpoint llamado “GetAll”.
4. Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.*/

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
