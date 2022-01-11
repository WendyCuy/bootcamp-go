/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del
sueldo y si gana más de $150.000 se le descontará además un 10%.*/

package main

import "fmt"

func main() {

	var salario float64 = 79000
	respuesta := impuestosalario(salario)
	fmt.Println("El valor del impuesto es: ", respuesta)

	var salario2 float64 = 190000
	respuesta2 := impuestosalario(salario2)
	fmt.Println("El valor del impuesto es: ", respuesta2)
}

func impuestosalario(sueldo float64) float64 {
	var impuesto float64
	if sueldo > 50000 && sueldo <= 149000 {
		impuesto = sueldo * 0.17
	} else if sueldo >= 150000 {
		impuesto = sueldo * 0.27
	}
	return impuesto
}
