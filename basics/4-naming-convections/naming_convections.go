package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	App       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	App       int
}


func main() {
	// PascalCase
	// e.g CalculateArea, UserInfo, NewHTTPRequest
	// STRUCTS, INTERFACES, enums

	// snake_case
	// e.g user_id, first_name, http_request

	// UPPERCASE
	// use case is Constants

	// mixedCase
	// e.g javasCript, htmlDocument, isValid

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)

}
