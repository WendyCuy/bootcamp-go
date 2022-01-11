package main

import (
    "fmt"
    "os"
)

type MyError struct{}

func (m *MyError) Error() string {
    if salary 
}

func main() {
    salary := 150000
	s, err := sayHello()
    if err != nil {
        fmt.Println("unexpected error: err:", err)
        os.Exit(1)
    }
    fmt.Println("The string:", s)
}