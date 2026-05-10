package main

import "fmt"

func main() {
	// go-lang only have keyword "for" for loops
	// for seperti while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for klasik: init; kondisi; post
	for i := 0; i < 5; i++ {
		fmt.Println(i, "masuk loop")
	}

	// for range
	names := []string{"Ikhsan", "rafi", "azhar"}

	// iterasi slice
	for index, value := range names {
		fmt.Println("index: , value: %s\n", index, value)
	}

	// iterasi map
	person := map[string]string{
		"name": "ikhsan",
		"age":  "23",
	}
	for key, value := range person {
		fmt.Println("key: %s, value: %s\n", key, value)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
