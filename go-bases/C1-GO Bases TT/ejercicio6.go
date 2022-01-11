package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])

	for key, element := range employees {
		if element > 21
	}

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)

}