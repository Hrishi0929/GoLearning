package main

// remember when we declare a package main we will always have to add a function main as well
// which will be executed when this program runs
// func main() {

// 	// the keyword string tells the go compiler that the value assigned to this variable will be a string
// 	// var card string = "Ace of Spades" // here we declare and assign a value to a variable called card
// 	// card := "Ace of Spades" // another way of declaring a variable // the := is used when the variable is initialized after that we can just use = to assign a variable a new value
// 	// card = "Five of Diamonds"
// 	card := newCard()

// 	fmt.Println(card)

// }

// func main() {
// 	// array is a fixed length list of things
// 	// slics is a fancy name for an array that can grow or shrink dynamically
// 	// every element in slices or arrays must be of the same type

// 	// slice declaration
// 	cards := []string{"Ace of Diamonds", newCard()}

// 	//adding a new element to the slice append(existing sliced, new element that we want to append to the end of the slice)
// 	cards = append(cards, "Six of Spades") // the append function doesn't modify the existing slice instead it returns a new slice that we then assign it back to the original slice in this case

// 	// how do we iterate over the slice cards
// 	// i --> index of the element in the array
// 	// card --> current card we are iterating over
// 	// range cards --> takes the slice of 'cards' and loops over it.
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}

// }

func main() {

	// since we have defined the type deck in deck.go which is a slice of strings
	// we can now replace []string to deck
	// cards := newDeck()

	// cards.print()
	// hand, remainingCards := deal(cards, 5) // since deal function returns 2 values so we initialize and assign the first value returned to hand variable and initialize and assign the second returned value to remainingCards variable
	// hand.print()
	// remainingCards.print()

	// here i have again initialized cards variable because we are using a new method called toString
	// we store the initialized deck to a file called myCards
	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("myCards")

	// here we initialize cards again be cause we ewant to read from a file that we stored
	// cards := newDeckFromFile("myCardsss")
	// cards.print()

	// we are re initializing cards again here to check how the shuffle function works
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string { // we need to specify the type of the return value after the paranthesis whenever newCard() is called it will now return a type
// 	return "Five of Diamonds"
// }
