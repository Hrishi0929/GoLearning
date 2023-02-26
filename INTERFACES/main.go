package main

import (
	"fmt"
)

// this is how we initialize an interface
// refer to the link below which will help understand why we used interfaces here and explains what interfaces do
// https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797378#overview
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{} // this is how we declare a new value of type struct when there's no value associated with it

	printGreeting(eb)
	printGreeting(sb)
}

// here we can see that we can call both the receiver functions that return english and spanish
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(engBot englishBot) {
// 	fmt.Println(engBot.getGreeting())
// }

//	func printGreeting(spaBot spanishBot) {
//		fmt.Println(spaBot.getGreeting())
//	}
func (engBot englishBot) getGreeting() string {
	// very custom logic to generate english greeting
	return "Hi There!!"
}

func (spaBot spanishBot) getGreeting() string {
	return "Hola!!"
}
