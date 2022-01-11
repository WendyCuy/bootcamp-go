package calculadora

import "testing"

func TestSumar(t *testing.T) {
	//preparación:  se dejan todas las variables o estructuras para después ejecutar
	a := 5
	b := 3
	resultadoEsperado := 8

	//acción:  se lleva a cabo la ejecución de la función a probar
	resultado := Sumar(a, b)

	//verificación de resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma() arrojó el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)

	}
}

func TestRestar(t *testing.T) {
	//preparación:  se dejan todas las variables o estructuras para después ejecutar
	a := 8
	b := 3
	resultadoEsperado := 5

	//acción:  se lleva a cabo la ejecución de la función a probar
	resultado := Restar(a, b)

	//verificación de resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion resta() arrojó el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)

	}
}
