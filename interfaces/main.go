package main

import "fmt"

type (
	bot interface {
		getGreeting() string
	}
	englishBot struct{}
	spanishBot struct{}
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}

// can remove sb since receiver var never used
func (spanishBot) getGreeting() string {
	return "Hola!"
}
