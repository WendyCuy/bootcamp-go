package ordenamiento

import "testing"

func TestOrdenar(t *testing.T) {
	//preparación:  se dejan todas las variables o estructuras para después ejecutar
	numeros := []int{2, 3, 1}
	resultadoEsperado := []int{1, 2, 3}

	//acción:  se lleva a cabo la ejecución de la función a probar
	resultado := Ordenar(numeros)

	//verificación de resultados
	validador := stringSlicesEqual(resultado, resultadoEsperado)
	if validador != true {
		t.Errorf("Funcion ordenar() arrojó el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)

	}

}

func stringSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
