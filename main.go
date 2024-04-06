package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) // Create a byte slice with 99999 elements
	// resp.Body.Read(bs)        // Read the response body into the byte slice

	// fmt.Println(string(bs)) // Print the response body

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

// go through the video 67, 68, 69, 70, 71 of Udemy -> Go: The Complete Developer's Guide (Golang)

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("asasa", len(bs))
	return len(bs), nil
}
