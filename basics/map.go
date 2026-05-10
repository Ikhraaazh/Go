package main

import "fmt"

func main() {
	// creating map
	var person map[string]string = map[string]string{}
	person["name"] = "ikhsan"
	person["address"] = "Jakarta"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	person["name"] = "ucup"
	fmt.Println(person["name"])

	delete(person, "name")
	fmt.Println(person)

	// more simple declaration
	orang := map[string]string{
		"name":    "ikhsan",
		"address": "yogya",
	}

	fmt.Println(len(orang))
}
