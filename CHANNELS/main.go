package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.org",
	}

	//creating a channel so that the main go routine can communicate with the child go routines and exit after the child go routines are done executing
	// we call the channel 'c' bloody 😡 go naming conventions
	// we use make() to intialize the channel that is a built in function that will create a value out of the given type
	c := make(chan string)

	for _, website := range websites {
		// checkIfWebsiteIsActive(website)

		go checkIfWebsiteIsActive(website, c) // we added the keyword go in front of checkIfWebsiteIsActive this makes a new go routine and makes a new thread for this function to run
		// we can only add go keyword in front of functions
	}

	// this is how we receive the message from our channel and log it
	// fmt.Println(<-c)

	// the problem that we are facing now is that when the first child routine sends a message on the channel
	// the main routine receives that message up and then exits the program thus leaving the other child routines hanging
	// to take care of that we need to write a for loop that will take care of that problem of main routine exiting
	// which makes the fmt.Println(<-c) piece of code a blocking call that's why we need to receive the message from the channel multiple times

	// for i := 0; i < len(websites); i++ { // notice that this is the traditional syntax of the for loop that we use
	// 	fmt.Println(<-c)
	// }

	// this is a loop that will run all the time till we don't interupt the code
	// so we are receiving the site in the channel message which we are then using in the checkIfWebsiteIsActive function
	// for {
	// 	go checkIfWebsiteIsActive(<-c, c)
	// }
	// alternate syntax for the above for loop

	// we are gonna run through the for loop every single time the channel emits a message or value
	// for link := range c {
	// 	go checkIfWebsiteIsActive(link, c)
	// }

	// we use a function literal here which is nothing but an unnamed function
	// that we use to wrap some little chunk of code that we can execute sometime in the future
	for link := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkIfWebsiteIsActive(link, c)
			// refer to this to understand why we need to pass link in the function literal
			// https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7824514#overview
		}(link) // the paranthesis is needed necause we need to execute the function literal here
	}
}

// function declaration without channel

// func checkIfWebsiteIsActive(site string) {
// 	_, err := http.Get(site)
// 	if err != nil {
// 		fmt.Println(site, "might be down...😩")
// 		return
// 	}
// 	fmt.Println(site, "is up and running...🥳")

// }

// since we want to use the channel to communicate between the go child routines and the go main routine
// we pass the channel c of type chan to the function checkIfWebsiteIsActive and we have to pass in string as well
// since we decided that the data shared via the channel is going to be of type string

// function declaration with channel
func checkIfWebsiteIsActive(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "might be down...😩")
		// this is how we send the message into our channel
		// c <- "Might be down I think"
		// let's push the site argument that we are receiving into the channel
		c <- site
		return
	}
	fmt.Println(site, "is up and running...🥳")
	// this is how we send the message into our channel
	// c <- "Yup it's up..."
	// let's push the site argument that we are receiving into the channel
	c <- site

}
