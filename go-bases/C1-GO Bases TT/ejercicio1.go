/*La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada 
una de las letras por separado para deletrearla.
1. Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de
letras que tiene la misma.
2. Luego imprimí cada una de las letras.*/

package main

import "fmt"

func main() {
	palabra := "Hola"

	fmt.Println("La cantidad de letras en la palabra", palabra, "es : ", len(palabra))

	for i, c := range palabra {
		fmt.Printf("%d: %s\n", i, string(c))
	}
}