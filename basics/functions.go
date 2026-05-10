package main

import "fmt"

// BASE
func sayHello() {
	fmt.Println("Hello")
}

// WITH PARAMETER
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// WITH RETURN VALUE
func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// WITH MULTIPLE RETURN VALUE
func getMinMax(array []int) (int, int) {
	min := array[0]
	max := array[0]

	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

// WITH NAMED RETURN VALUES
// with named return values, the declaration of the return values is done at the function declaration
func devide(a, b int) (result int, reminder int) {
	result = a / b
	reminder = a % b

	// just leave the return as is, no need to return anything here
	return
}

// VARIADIC FUNCTION
// the last parameter can hold more than 1 value (like array), notation: ...DataType
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// FUNCTION AS PARAMETER
// make alias for function type
type Transformer func(int) int

func printResult(value int, transformer Transformer) {
	result := transformer(value)
	fmt.Println("Result:", result)
}

func double(value int) int {
	return value * 2
}

func square(value int) int {
	return value * value
}

func main() {
	// base
	sayHello()

	// with parameter
	sayHelloTo("ikhsan", "rafi")

	// with return value
	fullName := getFullName("ikhsan", "azhar")
	fmt.Println(fullName)

	// with multiple return value
	numbers := []int{1, 2, 3, 4, 5}

	min, _ := getMinMax(numbers) // ignore max
	fmt.Println("Min:", min)

	_, max := getMinMax(numbers) // ignore min
	fmt.Println("Max:", max)

	// named return values
	result, reminder := devide(10, 2)
	fmt.Println("Result:", result)
	fmt.Println("Reminder:", reminder)

	// variadic function
	fmt.Println(sum(1, 2, 3))

	// slice as variadic function
	total := sum(numbers...)
	fmt.Println(total)

	// store function inside a variable
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(3, 20))

	// function as parameter
	printResult(10, double)
	printResult(10, square)

	// using anonymous function
	printResult(15, func(i int) int { return i * i })

}
