package main

import "fmt"

func main() {
	// var abcd map[string]string
	// abcd := make(map[string]int) print -> map[]
	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#4bf745",
		"orange": "#ff4500",
		"purple": "#800080",
		"black":  "#000000",
		"aqua":   "#00ffff",
	}
	colors["white"] = "#ffffff"
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}

/*
	Maps                        					| Struct
	------------------------------------------------|-------------------------------------------------
	- All keys & values must be the same  			| - Values can be different types
	- All keys are indexed, and we can iterate		| - Keys dont support indexing
	- Reference type								| - Value type
	- Use to represent a collection of related		| - Use to represent a "thing" with a lot of different properties
		values
	- Don't need to know all the keys at compile	| - Need to know all the fields at compile time
		time
	- Use when you need to look up a value by a		| - Use when you need to represent a "thing" with a lot of different properties
		single key
	- Use when you need a collection of related		| - Use when you need to represent a fixed collection of fields
		values
	- Use when you need a key-value pair			| - Use when you need a collection of fields
	- cant use . operator to access the value		| - can use . operator to access the value]
*/
