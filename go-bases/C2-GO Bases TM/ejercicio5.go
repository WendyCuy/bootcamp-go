/* Ejercicio 5 - Calcular cantidad de alimento
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
haber muchos más animales que refugiar.
Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.
Se solicita:
- Implementar una función Animal que reciba como parámetro un valor de tipo texto
con el animal especificado y que retorne una función y un error (en caso que no exista
el animal)
- Una función para cada animal que calcule la cantidad de alimento en base a la
cantidad del tipo de animal especificado.*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	animalPerro, err := animal(perro)
	animalGato, err := animal(gato)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	if err != nil {
		fmt.Println(err)
	}

	var cantidad float64
	cantidad += animalPerro(10)
	cantidad += animalGato(10)
	cantidad += animalHamster(5)
	cantidad += animalTarantula(8)

	fmt.Println("Cantidad de alimento que se debe comprar:", cantidad, "Kg")
}

func perro(cantidadPerros float64) float64 {
	total := cantidadPerros * 10
	return total
}

func gato(cantidadGatos float64) float64 {
	total := cantidadGatos * 5
	return total
}

func hamster(cantidadHamsters float64) float64 {
	total := (cantidadHamsters / 1000) * 250
	return total
}

func tarantula(cantidadTarantulas float64) float64 {
	total := (cantidadTarantulas / 1000) * 150
	return total
}

func animal(animal string) (func(float64) float64, error) {

	switch animal {
	case "perro":
		funcionPerro := perro
		return funcionPerro, nil
	case "gato":
		funcionGato := gato
		return funcionGato, nil
	case "hamster":
		funcionHamster := hamster
		return funcionHamster, nil
	case "tarantula":
		funcionTarantula := tarantula
		return funcionTarantula, nil
	default:
		return nil, errors.New("El animal no existe")
	}

}