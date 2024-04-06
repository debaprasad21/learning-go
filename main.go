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

	// it has the address of the value of jim
	// jim is a reference to the struct in memory the actual value of the struct
	// Pointer Shortcut
	jim.updateName("Jimmy")

	jim.print()

	myArrays := [5]string{"Hi", "There", "How", "Are", "You"}
	updateArray(myArrays)
	fmt.Println(myArrays)

	mySlices := [5]string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlices)
	fmt.Println(mySlices)

	name := "Bill"
	fmt.Println(*&name)

}

/*
	value types:		|   reference types:
	--------------------| ---------------------
	int					|	slices
	float				|	maps
	string				|	channels
	bool				|	pointers
	struct				|	functions

*/

// in slice, we don't need to use * to pass by reference
// because they are already references
func updateSlice(s [5]string) {
	s[0] = "Bye"
}

func updateArray(s [5]string) {
	s[0] = "Bye"
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

/*
	jimPointer := &jim
	jimPointer.updateName("Jimmy") -> jimPointer is Type of *person, or a pointer to a person

	jim.updateName("Jimmy") -> jim is Type of person, or a person

	func (pointerToPerson *person) updateName() -> *person is a type of *person, it means we are working with a pointer to a person
*/
