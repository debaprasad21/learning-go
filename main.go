package main

import (
	"fmt"
	"strings"
)

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available and %v tickets remaining\n", conferenceTickets, remainingTickets,)
	fmt.Println("Get your tickets to attend")

	for {
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
		bookings = append(bookings, fName + " " + lName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fName, lName, userTickets, email)
		fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)
		
		fNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			fNames = append(fNames, names[0])

		}
		fmt.Printf("The first names of the bookings are: %v\n", fNames)	
	}
}