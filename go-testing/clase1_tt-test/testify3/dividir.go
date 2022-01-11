package dividir

import "errors"

func Division(a, b int) (int, errors) {
	num := int
	den := int

	if den == 0 {
		return 0, errors.New("El denominador no puede ser cero")
	} else {
		return num / den, nil
	}
}
