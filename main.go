package main

import (
	"fmt"
	"sync"
	"time"
)

// Define variables as possible, only when the variables are common to all can be defined at package level

// package variable, scope - package

// Define a struct for user data
// type keyword is used to define a new data type
// struct is a collection of fields
type UserData struct {
	fName       string
	lName       string
	email       string
	noOfTickets uint
}

// data type is slice of struct - having list of struct defined in UserData
// here 0 represents the length of the slice and increases as we add more elements
const conferenceTickets int = 50

var bookings = make([]UserData, 0)
var conferenceName = "Go Conference"
var remainingTickets uint = 50

// waits for the launched goroutine to finish
// Package sync provides basic synchronization primitives such as mutual exclusion locks
// Add: Sets the number of goroutines to wait for (increases the counter by the provided number)
// Wait: Blocks until the counter is zero
// Done: Decreases the counter by one, So this is called by the goroutine to indicate that it has finished
var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	fName, lName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidUserTickets := validateUserInput(fName, lName, email, userTickets)

	if isValidName && isValidEmail && isValidUserTickets {

		bookTicket(userTickets, fName, lName, email)
		// to handle the concurrency
		wg.Add(1)
		go sendTicket(userTickets, fName, lName, email)

		fmt.Printf("The first names of the bookings are: %v\n", getFirstNames())

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			fmt.Println("All tickets are sold out")
			// ending the infinite loop
			// break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets available and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames() []string {
	fNames := []string{}
	for _, booking := range bookings {
		fNames = append(fNames, booking.fName)

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

	var userData = UserData{
		fName:       fName,
		lName:       lName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", fName, lName, userTickets, email)
	fmt.Printf("Remaining tickets for %v are %v\n", conferenceName, remainingTickets)
}

func sendTicket(userTickets uint, fName string, lName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, fName, lName)
	fmt.Println("----------------------")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("----------------------")
	wg.Done()
}
