package main

import "fmt"

func main() {

	n := 3

	_, msg := isPrime(n)
	fmt.Println(msg)

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
