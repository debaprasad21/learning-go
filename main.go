package main

import (
	"fmt"
	"strconv"
)

// Define variables as possible, only when the variables are common to all can be defined at package level

// package variable, scope - package

const conferenceTickets int = 50

var bookings = make([]map[string]string, 0) // data type is slice of map - having list of maps
// here 0 represents the length of the slice and increases as we add more elements
var conferenceName = "Go Conference"
var remainingTickets uint = 50

func main() {
	greetUsers()

	for {
		fName, lName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(fName, lName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, fName, lName, email)

			fmt.Printf("The first names of the bookings are: %v\n", getFirstNames())

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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames() []string {
	fNames := []string{}
	for _, booking := range bookings {
		fNames = append(fNames, booking["firstName"])

	}
	return fNames
}

func getUserInput() (string, string, string, uint) {
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

	return fName, lName, email, userTickets
}

func bookTicket(userTickets uint, fName string, lName string, email string) {
	remainingTickets -= userTickets

	// create a map for the user
	// make - to create empty map
	// all keys have the same data type
	// all values have the same data type
	var userData = make(map[string]string)

	userData["firstName"] = fName
	userData["lastName"] = lName
	userData["email"] = email
	userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// converting userTickets to string since in map all keys and values are strings

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fName, lName, userTickets, email)
	fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)
}
