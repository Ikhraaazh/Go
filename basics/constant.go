package main

import "fmt"

func main() {
	const PI = 3.14
	const max_value = 100

	fmt.Println(PI)
	fmt.Println(max_value)

	// PI = 3.15 // ERROR: cannot assign to PI

	// multiple constant declaration
	const (
		value  = 100
		number = 20
	)

	fmt.Println(value, number)
}
