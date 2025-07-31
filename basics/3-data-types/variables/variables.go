package main

import "fmt"

// := symbol can only be used within functions

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	middleName := "Cane"
	// middleName := "Mayor"

	fmt.Println(middleName)

	// Default values
	// Numeric Types: 0
	// Boolean Types: false
	// String Types: ''
	// Pointers, slices, maps, functions and structs: nil

	// .... SCOPE


}

func printname() {
	firstName := "Micheal"
	fmt.Println(firstName)
}
