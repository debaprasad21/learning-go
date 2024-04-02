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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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

		isValidName := len(fName) > 2 && len(lName) > 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, fName+" "+lName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fName, lName, userTickets, email)
			fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)

			fNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				fNames = append(fNames, names[0])

			}
			fmt.Printf("The first names of the bookings are: %v\n", fNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("All tickets are sold out")
				// ending the infinite loop
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is invalid. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid. Please try again.")
			}
			if !isValidUserTickets {
				fmt.Println("No of ticket entered is invalid. Please try again.")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have %v tickets available and %v tickets remaining\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}
