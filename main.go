package main

import "fmt"

type bot interface {
	// mengambil fungsi getGreeting yang ada dibawah
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)

}

func (eb englishBot) getGreeting() string {
	//	custom logic in english
	return "hello"
}

func (sb spanishBot) getGreeting() string {
	//	custom logic in spanish
	return "hallo"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
