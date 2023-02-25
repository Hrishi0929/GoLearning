package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type called deck
//
//	which is a slice of strings
type deck []string

// this function will create and return a list of playing cards
func newDeck() deck { // this function will return a value of type deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"} // this can be expanded to go all the way to king depends on us

	// here we first loop through the cardSuits list and then
	// we loop through the list of cardValues and for every combinaiton
	// of suit and value we add a new card to the cards deck/slice/list

	// whenever we have a variable that we actually don't want to use like the i and j variables here
	// in go we can replace them with _ and go understands that there's a variable here but  we don't want to use these variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function sets up methods on vaiables we create
// d is reference to the instance of the deck variable that we are working with
// now  we can imagine that when print function is executed  on the cards variable
// the that cards variable gets passed into the print function as the variable that we have called d.

// by convention we usually refer to the receiver with a one or two letter abbreviation that matches
// the type of the receiver in this case our receiver is of type deck so we refer the reciver as 'd'

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// any variable of type deck now gets access to the print method
// because cards is of type deck now we can now call the print method on the cards variable

// convention of the arguments passed in the function is the name of the argument followed by the type of the argument
func deal(d deck, handSize int) (deck, deck) { // here we are returning 2 values of type deck so we have to add a pair of parenthesis after the func declaration and add the return types
	// slice1 := d[:handSize] //d[startIndexIncluding : upToNotIncluding]  d[:handSize] = d[0:handSize]
	// slice2 := d[handSize:] d[handSize:] = d[handSize : endOfThe Slice/List/Array]

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	// we can do this because we derived the type deck from the string so we can go from type deck to type string pretty easily by doing it like this
	return strings.Join([]string(d), ",") // strings.Join takes the string slice as the first arg and the seperator as the second arg

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is the permission for anyone can read and write this file

}

func newDeckFromFile(fileName string) deck {
	biteSlice, err := ioutil.ReadFile(fileName) // ioutil.Readfile returns two values one is the biteSlice and other is the error object
	if err != nil {
		// whenever error handling is being done in go first wait and take some time to understand what can go wrong
		// Option #1 --> log the error and return a call to newDeck()
		// Option #2 --> Log the error and entirely quit the program
		fmt.Println("Here is an error message😢 to review: ", err)
		os.Exit(1) // here if we pass anything other than 0 that means there is an error in the program and os.Exit will quit the program entirely
	}

	sliceOfStrings := strings.Split(string(biteSlice), ",") // the first arg is the string that we want to split and the second arg is the separator which is "," in our case
	return deck(sliceOfStrings)
}

func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // here we are taking whatever is at i and place it in the newPosition index and the element present at newPosition index to the index i
		// inturn generating a random shhuffled deck
	}
}
