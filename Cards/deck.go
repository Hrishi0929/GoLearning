package main

import "fmt"

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
