/* Ejercicio 2 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
necesitan una aplicación que les permita calcular el descuento con base en 2 variables, su
precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
descuento aplicado y luego imprimirlo en consola.
● Crear la aplicación de acuerdo con los requerimientos.*/

package main

import "fmt"

func main() {

	var precio int = 1000
	var porcent_descuento int = 10

	descuento := (precio /100)  * porcent_descuento
	valorpago := precio - descuento

	fmt.Println(" El valor a pagar es  $", valorpago, "el descuento es de: ", descuento, "sobre el valor de su compra")

}
