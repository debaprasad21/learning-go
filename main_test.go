package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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
// go test -run Test_isPrime -> to run a specific test
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
		{"zero", 0, false, "0 is not a prime number"},
		{"negative", -1, false, "-1 is negative number & is not a prime number"},
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

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to the write pipe
	os.Stdout = w

	prompt()

	// close the write pipe
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}

}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to the write pipe
	os.Stdout = w

	intro()

	// close the write pipe
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro test not correct but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"prime", "7", "7 is a prime number!"},
		{"not prime", "8", "8 is not a prime number because it is divisible by 2"},
		{"zero", "0", "0 is not a prime number"},
		{"negative", "-1", "-1 is negative number & is not a prime number"},
		{"invalid", "a", "Please enter a valid number"},
		{"quit", "q", "Goodbye!"},
	}

	for _, e := range tests {

		// input the number
		input := strings.NewReader(e.input)

		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function we need a channel and an instance of an io.Reader

	doneChan := make(chan bool)

	//create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan

	close(doneChan)

}
