package main

import "fmt"

func main() {
	// deklarasi arrray dengan ukuran 3
	var names [3]string
	names[0] = "ikhsan"
	names[1] = "rafi"
	names[2] = "asep"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(len(names))

	// inisialisasi langsung
	nama := [4]string{"ikhsan", "rafi", "asep", "agus"}

	// atau pake ... untuk auto-count
	hobbies := [...]string{"reading", "gaming"}

	fmt.Println(nama)
	fmt.Println(hobbies)
	fmt.Println(len(hobbies))

	// array[low:high] dari index low sampai sebelum high
	slice1 := nama[1:3] // ["rafi", "asep"]
	slice2 := nama[1:]  // ["rafi", "asep", "agus"]
	slice3 := nama[:3]  // ["ikhsan", "rafi", "asep"]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
