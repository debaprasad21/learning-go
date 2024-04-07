package main

import "testing"

// The file name should end with _test.go
// The function name should start with Test follewed by the function name in main.go not in small letters
// The function should accept a pointer to testing.T
// The function should be in the same package as the main code
// The function should be in the same directory as the main code
// go test -> to run the test
// go test -v -> to run the test in verbose mode and prints the results of the test
// go test -cover -> to check the code coverage
// go test -coverprofile=coverage.out -> to check the code coverage and save the results in a file, here the filename is coverage.out
// go tool cover -html=coverage.out -> to view the code coverage in a browser
func Test_isPrime(t *testing.T) {
	// creating a slice of struct to test the isPrime function
	// test table
	// entries to test various scenarios
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"not prime", 0, false, "0 is not a prime number"},
		{"not prime", -1, false, "-1 is negative number & is not a prime number"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
