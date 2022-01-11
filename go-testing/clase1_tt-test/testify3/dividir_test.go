package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 1

	// Se ejecuta el test
	resultado := Division(num1, num2)(error)

	// Se validan los resultados aprovechando testify
	assert.NotNil(t, resultado)
}
