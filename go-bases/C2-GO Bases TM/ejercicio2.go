/* Ejercicio 2 - Calcular promedio
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva
el promedio y un error en caso que uno de los números ingresados sea negativo */

package main

import (
	"errors"
	"fmt"
)

func sum(values ...int) (int, error) {
	var result int
	var promedio int

	for _, value := range values {
		if value < 0 {
			return 0, errors.New("No pueden existir números negativos")
		} else {
			result += value
		}
	}
	promedio = result / len(values)
	return promedio, nil

}
func main() {
	x, err := sum(1, 2, 3, 4, 5, 6, 10, 1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result :", x)
	}

}
