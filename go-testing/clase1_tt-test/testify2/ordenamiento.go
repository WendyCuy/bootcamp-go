package ordenamiento

import "sort"

func Ordenar(a []int) []int {
	sort.Ints(a)
	return a
}
