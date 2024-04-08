package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye!")

}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it's a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	// scanner := bufio.NewScanner(os.Stdin)

	// for {
	// 	res, done := checkNumbers(scanner)

	// 	if done {
	// 		doneChan <- true
	// 		return
	// 	}

	// 	fmt.Println(res)
	// 	prompt()
	// }

	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}

}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	// read user input
	scanner.Scan()

	// check if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "Goodbye!", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a valid number", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 & 1 are not prime numbers
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number", n)
	}

	// negative numbers are not prime numbers
	if n < 0 {
		return false, fmt.Sprintf("%d is negative number & is not a prime number", n)
	}

	// use the modulus operator to check if the number is prime
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
