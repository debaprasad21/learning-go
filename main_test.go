package main

import "testing"

// The file name should end with _test.go
// The function name should start with Test follewed by the function name in main.go not in small letters
// The function should accept a pointer to testing.T
// The function should be in the same package as the main code
// The function should be in the same directory as the main code
func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not a prime number" {
		t.Errorf("Wrong message returned:- %s", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 7)
	}

	if msg != "7 is a prime number!" {
		t.Errorf("Wrong message returned:- %s", msg)
	}
}
