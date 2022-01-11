/* Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido,
DNI, Fecha y que tenga un método detalle */

package main

import "fmt"

type alumnos struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (v alumnos) detalle() {

	fmt.Printf("\nnombre: %s \napellido: %s \ndni: %v \nfecha: %s\n", v.nombre, v.apellido, v.dni, v.fecha)
}

func main() {
	p1 := alumnos{
		nombre:   "Hernan",
		apellido: "Cuy",
		dni:      "1030",
		fecha:    "13-12-2021",
	}

	p2 := alumnos{
		nombre:   "Luz",
		apellido: "Salcedo",
		dni:      "6574",
		fecha:    "13-12-2021",
	}

	p3 := alumnos{
		nombre:   "David",
		apellido: "Cuy",
		dni:      "9336",
		fecha:    "13-12-2021",
	}
	p1.detalle()
	p2.detalle()
	p3.detalle()
}
