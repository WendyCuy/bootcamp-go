/* Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
por mes y la categoría, y que devuelva su salario.*/

package main

import "fmt"

func main() {
	var (
		minutosTrabajados float64 = 10000
		categoriaEmpleado string  = "C"
	)
	respuesta := salariomes(minutosTrabajados, categoriaEmpleado)
	fmt.Println("El salario del mes es de ", respuesta)
}

func salariomes(minutos float64, categoria string) float64 {
	var (
		hora          float64 = 60
		salario       float64
		horaTrabajada = minutos / hora
	)

	switch categoria {
	case "C":
		salario = horaTrabajada * 1000
		return salario
	case "B":
		salario = (horaTrabajada * 1500) + ((horaTrabajada * 1500) * 0.20)
		return salario
	case "A":
		salario = (horaTrabajada * 3000) + ((horaTrabajada * 3000) * 0.50)
		return salario
	}
	return 0
}
