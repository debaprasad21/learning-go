package main

import "fmt"

type contactInfo struct {
	email        string
	zipCode      int
	mobileNumber int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// ways to initialize struct
	// alex := person{"Alex", "Anderson"}
	// best practice to use key value pair
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:        "jim@gmail.com",
			zipCode:      94000,
			mobileNumber: 1234567890,
		},
	}

	// jimPointer corresponds to the memory address of jim,
	// it has the address of the value of jim
	// jim is a reference to the struct in memory the actual value of the struct
	jimPointer := &jim
	jimPointer.updateName("Jimmy")

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// *person is a type description, it means we are working with a pointer to a person
// *pointerToPerson is an operator, it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson is the actual value of the memory address
	// (*pointerToPerson).firstName is the value of the firstName of the actual value of the memory address
	// () is used to make sure that the . operator is used before the * operator
	fmt.Printf("Memory address of the pointer: %v & actual value is: %v\n", pointerToPerson, *pointerToPerson)
	(*pointerToPerson).firstName = newFirstName
}

/*
	-------------------------------------------------------------------------------------
	|						 	|														|
	| address -> 000xc0000b6010 |  value -> {firstName: "Jim", lastName: "Party" ....}	|
	|						 	|														|
	-------------------------------------------------------------------------------------
	* Turn address into value with *address
	* Turn value into address with &value
*/
