package main

// bot is an interface that has a method getGreeting
// getGreeting returns a string
// englishBot and spanishBot are structs that implement the bot interface
// printGreeting is a function that takes a bot interface and prints the greeting
type bot interface {
	// if you are a type in this program with a function called getGreeting and
	// you return a string you are now an honorary member of type bot
	getGreeting() string
}

// we use interfaces to define a contract between different types

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

/*
	Interfaces are not generic types
	Interfaces are 'implicit'
	Interfaces are a contract to help us manage types
	Interfaces are tough, but very powerful
	Interfaces are a way to make our code better

	example:
		type bot interface {
			getGreeting(string, int) (string, error)
		}
		bot -> interface name
		getGreeting -> method name
		(string, int) -> list of arguments
		(string, error) -> list of return types
*/
