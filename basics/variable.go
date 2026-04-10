package main

import "fmt"

func main() {
	// declaring variable with data type
	var nama string
	nama = "ucup"

	// declaration + initialization (type can be omitted)
	var umur = "25"

	// shorthand (without var, directly :=)
	alamat := "Rumah"

	fmt.Println(nama, umur, alamat)

	// multiple variable declaration
	var (
		firstName = "bagus"
		lastName  = "nurbagus"
	)

	fmt.Println(firstName, lastName)
}
