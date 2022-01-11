package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	//preparación:  se dejan todas las variables o estructuras para después ejecutar
	numeros := []int{2, 3, 1}
	resultadoEsperado := []int{1, 2, 3}

	//acción:  se lleva a cabo la ejecución de la función a probar
	resultado := Ordenar(numeros)

	//verificación de resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
