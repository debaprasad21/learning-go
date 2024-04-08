package main

import "fmt"

// we can debug the code by using fmt.Printf() to print the type of the variables
// or adding log statements to see the flow of the code
// logger is a package that is used to log the messages
// delve is a debugger that is used to debug the code
// delve link -> https://github.com/go-delve/delve/blob/master/Documentation/installation/osx/install.md
// delve commands -> dlv
// cmd+shift+p => mac command to open command palette
// cntrl+shift+p => windows command to open command palette
// Run and Debug in VSCode and create a launch.json file
/* Update the program  as given below to debug from workspace
 {
  // Use IntelliSense to learn about possible attributes.
  // Hover to view descriptions of existing attributes.
  // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}"
    }
  ]
}
*/
// add breakpoints in the code by clicking on the left side of the line number
// if u want to add breakpoints in the code by using command line use the below command
// break ./main.go:15 -> 15 is the line number
// ensure go.mod file is present in the project
// launch the debugger from Run and Debug in VSCode

func main() {
	count := 0

	for {
		if count == 10 {
			break
		}
		count++

	}
	fmt.Println("Counter", count)

}
