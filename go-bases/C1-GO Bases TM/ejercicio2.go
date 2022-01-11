/* Ejercicio 2 - Clima

Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura, humedad y 
presión atmosférica de distintos lugares. 
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y 
presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?*/

package main

import "fmt"

func main() {
	var (
		temperatura float64 = 14.0
		humedad     float64 = 100.0
		presion     int     = 1026
	)
	fmt.Println("Bogotá")
	fmt.Printf("La temperatura es %f°C y la variable es de tipo %T\n", temperatura, temperatura)
	fmt.Printf("La humedad relativa es %f %% y la variable es de tipo %T\n", humedad, humedad)
	fmt.Printf("La presion atmosférica es %d hPa y la variable es de tipo %T\n", presion, presion)
}

