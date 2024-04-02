package main

import (
	"strings"
)

func validateUserInput(fName string, lName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(fName) > 2 && len(lName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	// In Go we can return any number of values from a function
	return isValidName, isValidEmail, isValidUserTickets
}
