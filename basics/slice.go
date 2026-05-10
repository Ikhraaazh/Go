package main

import "fmt"

func main() {
	// append
	var hobbies []string // empty slice

	hobbies = append(hobbies, "gaming") // adding to slice
	hobbies = append(hobbies, "reading", "sleeping")

	fmt.Println(hobbies)      // [gaming, reading, sleeping]
	fmt.Println(len(hobbies)) // 3
	fmt.Println(cap(hobbies))

	// make slice
	slice := make([]string, 3, 10) // make(sliceType, length, capacity) // length 3, capacity 10
	slice[0] = "ikhsan"
	slice[1] = "rafi"
	slice[2] = "dzulfikar"

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// copy slice
	source := []string{"Ikhsan", "rafi", "azhar"}
	destination := make([]string, len(source))

	copy(destination, source)

	fmt.Println(source)
	fmt.Println(destination)

	// ARRAY VS SLICE DIFFERENCE
	array := [3]string{"a", "b", "c"}    // fixed size
	sliceExmp := []string{"a", "b", "c"} // dynamic size

	fmt.Println(array)
	fmt.Println(sliceExmp)
}
