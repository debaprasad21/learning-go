package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available and %v tickets remaining\n", conferenceTickets, remainingTickets,)
	fmt.Println("Get your tickets to attend")

	// when we know the values of the array
	// var bookings = [50]string{"John", "Deba"}

	// when we don't know the values of the array
	// var bookings = [50]string{}

	// alternate way to declare an array
	var bookings [50]string

	var fName string
	var lName string
	var email string
	var userTickets uint
	// ask user to enter their name
	fmt.Println("Enter first name:")
	fmt.Scan(&fName)

	fmt.Println("Enter last name:")
	fmt.Scan(&lName)

	fmt.Println("Enter email name:")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings[0] = fName + " " + lName

	fmt.Printf("The Whole Array: %v\n", bookings)
	fmt.Printf("The first value of the Array: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fName, lName, userTickets, email)
	fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)	

}