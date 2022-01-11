/* Ejercicio 3 - Préstamo
Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es menor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de
acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.*/

package main

import "fmt"

func main() {
	var (
		edad       int     = 19
		antiguedad int     = 5
		empleado   bool    = true
		sueldo     float32 = 120000
	)
	if edad > 22 {
		if empleado == true {
			if antiguedad > 1 {
				if sueldo < 100000 {
					fmt.Println("Genial! tu credito no generará intereses")
				} else {
					fmt.Println("Tu credito generará intereses")
				}
			} else {
				fmt.Println("Debes tener mas de un año de antiguedad en tu trabajo")
			}
		} else {
			fmt.Println("Lamentablemente debe tener empleo ")
		}
	} else {
		fmt.Println("Debe ser mayor a 22 años")
	}
}
