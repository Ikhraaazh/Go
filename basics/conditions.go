package main

import "fmt"

func main() {
	nilai := 70

	if nilai >= 80 {
		fmt.Println("A")
	} else if nilai >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("Tidak lulus")
	}

	// if else with short statement
	if nilai := 80; nilai >= 75 {
		fmt.Print("Lulus dengan nilai ", nilai)
	} else {
		fmt.Println("Tidak lulus dengan nilai ", nilai)
	}

	// switch
	scire := "B"

	switch scire {
	case "A":
		fmt.Println("Sangat bagus")
	case "B":
		fmt.Println("Bagus")
	default:
		fmt.Println("Tidak lulus")
	}

	switch {
	case nilai >= 90:
		fmt.Println("A")
	case nilai >= 80:
		fmt.Println("B")
	case nilai >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

}
