/* Ejercicio 4 - A qué mes corresponde
Realizar una aplicación que contenga una variable con el número del mes.
1. Según el número, imprimir el mes que corresponda en texto.
2. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por
qué?
Ej: 7, Julio*/


package main

import "fmt"

func main() {
	mes := 1
	switch mes {
	case 1:
		fmt.Println(mes, ", enero")
	case 2:
		fmt.Println(mes, ", febrero")
	case 3:
		fmt.Println(mes, ", marzo")
	case 4:
		fmt.Println(mes, ", abril")
	case 5:
		fmt.Println(mes, ", mayo")
	case 6:
		fmt.Println(mes, ", junio")
	case 7:
		fmt.Println(mes, ", julio")
	case 8:
		fmt.Println(mes, ", agosto")
	case 9:
		fmt.Println(mes, ", septiembre")
	case 10:
		fmt.Println(mes, ", octubre")
	case 11:
		fmt.Println(mes, ", noviembre")
	case 12:
		fmt.Println(mes, ", diciembre")
	default:
		fmt.Println("Desconocido")
	}
}