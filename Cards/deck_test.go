package main

import (
	"os"
	"testing"
)

// So I want to figure out how are we going to decide what to test in this case?
// Let's think about our function, new deck.
// Now, the new deck function has a decent amount of logic inside of it, right?
// Like if we look at the new deck function inside of our code editor.
// So here's new deck right here We create an empty deck.
// We put together these two kind of starting lists of suits and values, and then we've got this nested
// if statement excuse me, nested for statement.
// And so I can very easily imagine many situations where another developer might come in here and they
// start messing around with stuff and start breaking code accidentally.
// And so I think that it's completely reasonable to say maybe we want to write a test around this new
// deck thing.
// But even then I still get the question, well, what are we actually testing?
// Like, how do we write a test to make sure that new deck is working the way we expect?
// And my answer to you and the answer to all these test questions is that it's up to you as the developer
// to just say, What do I really care about with New Deck?
// What do you really care about?
// Well, for me personally, when I think about the new deck function, there's probably three things
// that I really care about, or three easy things I could very easily check.
// And if each of these little criteria are correct, when I run my tests, it probably means that new
// deck is working the way I expect.
// And so here's the three criteria.
// First off, I think that it makes sense to say that after we make a new deck, we should have, say,
// four items inside of it.
// So let's assume that we're making a deck with four items or four strings, four cards.
// So maybe it makes a lot of sense to say, let's check the length of the deck that gets returned and
// it should be exactly like four cards.
// I think that's a completely reasonable check Right.
// Maybe the another thing I would want to check is to make sure that the very first card in the deck is
// equal to ace of spades.
// And then maybe a third thing that's a really easy and straightforward check is to say that I want to
// make sure that the last card is equal to the four of spades.
// And so there's really not a lot of trickery here or a lot of ambiguity.
// Whenever you write tests, just think about, well, what are some easy assertions?
// Or What do I kind of care about with the code that I'm writing?
// You can write some tests around that, and honestly, that's kind of all there is to it.
// So let's take those three tests and figure out exactly how we might structure them inside of our test file.

func TestNewDeck(t *testing.T) { //t is the test handler if something goes wrong then we will have to tell 't' what went wrong
	// we make a function called as func TestNewDeck --> to test the newDeck function
	// Code to make sure that a deck is created with x number of cards
	// Code to make sure that the first card is an Ace of Spades
	// Code to make sure that the last card is a FOur of Clubs

	// basically create a new deck
	// write an if statemente to see if the deck has the right number of cards
	// if it doesn't, tell the go test handler that something went wrong

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d)) // %v will take the value of len(d)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}

}

// Testing saveToDeck and newDeckFromFile
// Delete any files in current working directory with the name "_decktesting"
// Create a deck
// Save to file "_decktesting"
// Load from file
// Assert deck length
// Delete any files in current working directory with the name "_decktesting"
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
