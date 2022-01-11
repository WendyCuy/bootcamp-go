/* Ejercicio 5 - Listado de nombres
a. Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es
necesario crear una aplicación que contenga dicha lista.
Estudiantes:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán,
Leandro, Eduardo, Duvraschka.
b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
sin modificar el código que escribiste inicialmente.
Estudiante:
Gabriela*/

package main

import "fmt"

func main() {
	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán",
		"Leandro", "Eduardo", "Duvraschka"}

	estudiantes = append(estudiantes, "Gabriela")

	fmt.Println(estudiantes)
}