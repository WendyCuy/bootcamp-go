package main

import (
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

/* Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
2. Luego genera la lógica de filtrado de nuestro array.
3. Devolver por el endpoint el array filtrado.*/

func GenerarListaProductos([]string) {
	// consultar el Json y  guardarlos en el arreglo
}

/*func FiltrarProductos(ctxt *gin.Context) {
	productos := GenerarListaProductos()
	var filtrados []*Productos
	for i, e := range productos {
		if ctxt.Query("nombre") == e.Nombre {
			filtrados = append(filtrados, &e)
		}
	}
}
*/

/* Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
1. Genera una nueva ruta.
2. Genera un handler para la ruta creada.
3. Dentro del handler busca el item que necesitas.
4. Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.*/

func BuscarProducto(ctxt *gin.Context) {
	producto, ok := Productos[ctxt.Param("id")]
	if ok {
		ctxt.String(200, "Información del producto %s, producto: %s", ctxt.Param("id"), producto)

	} else {
		ctxt.String(404, "Información del producto ¡No existe!")
	}
}

func main() {
	server := gin.Default()
	server.GET("/products/:id", BuscarProducto) // Genera una nueva ruta y el handler.
	server.Run(":8080")
}
