package main

import "fmt"

// whenever we don't assign an actual value to a variable go assigns it with the zero value
// for a map the zero value is an empty map
func main() {
	//1st way to declare a map
	colors := map[string]string{ //here we are declaring a map where all of the keys are of type string and all of the values are also of type string
		"red":   "#FF0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// fmt.Println(colors)

	//2nd way
	// var colors map[string]string
	// fmt.Println(colors)

	//3rd way
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff" // this is how we add data to an empty map or add more data to an already initialized map

	// delete(colors, "white") // this is how we delete an entry from the map we pass the map variable and then pass the key of the entry we want to delete
	// fmt.Println(colors)
	printMap(colors)

}

// iterating over a map
func printMap(mapArgument map[string]string) {
	for color, hex := range mapArgument { // here color and hex are the key and value of our colors map so
		// when we run a for loop on a map we will iterate on the key and value instead of the index and value that we do normally
		fmt.Println("Hex code for", color, "is", hex)

	}
}

// Differences between maps and structs
//https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797370#overview
// refer the above link to get a clearer picture
