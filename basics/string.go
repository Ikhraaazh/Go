package main

import "fmt"

func main() {
	var nama string = "ikhsan rafi"
	fmt.Println(nama)
	fmt.Println(len(nama))       // string lenght
	fmt.Println(nama[0])         // first character (in byte)
	fmt.Println(string(nama[0])) // "i"
}
