package main

import "fmt"

func main() {
	age := 32           // regular variable
	var agePointer *int // pointer to the variable
	agePointer = &age   // assigning address of age to the pointer

	fmt.Println("Value of age:", age)
	fmt.Println("Address of age:", agePointer) // address of the variable

	editAgeToAdultYears(agePointer)
	fmt.Println("Years as an adult:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18 // dereferencing the pointer to get the value
	*age = *age - 18
}
