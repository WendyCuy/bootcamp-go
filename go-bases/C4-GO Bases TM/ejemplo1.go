package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myCustomError{
			status: status,
			msg:    "algo salió mal",
		}
	}
	return 200, nil
}

func main() {
	status, err := myCustomErrorTest(300)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Status %d, Funcional", status)
}