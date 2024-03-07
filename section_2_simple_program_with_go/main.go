package main

import "fmt"

func main() {
	var name = "Taufik"
	age := 34
	isMentor := false

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Is mentor: ", isMentor)

	var radius float64 = 5.0
	var area float64
	const pi = 3.14

	area = pi * (radius * radius)
	fmt.Println("Area: ", area)
}
