/* Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma (csv).
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.*/

package main

import (
	"io/ioutil"
	"log"
)


func main(){
	producto := "{id: 1, nombre: shampoo, precio: $100, cantidad: 10};{id: 2, nombre: cepillo, precio: $50, cantidad: 12};{id: 2, nombre: jabon, precio: $75, cantidad: 21}"
	archivo := []byte(producto)
	err := ioutil.WriteFile("./listadoProductos.txt", archivo, 0644)
	if err != nil {
		log.Fatal(err)
	}
}