// Package deck provides various deck capabilities
package deck

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Deck represents an individual deck of cards
type Deck struct {
	Name  string   `json:"name"`
	Cards []string `json:"cards"`
}

// ErrCardNotFound is an error when a user asks for a card not found within a deck.
var ErrCardNotFound = errors.New("card not found with that index")

// ModifiedFibonacci is the standard deck for agile estimations.
var ModifiedFibonacci = &Deck{"Modified Fibonacci", []string{"0", "1", "2", "3", "5", "8", "13", "20", "40", "100", "?", "☕"}}

// Fibonacci uses the actual Fibonacci numbers.
var Fibonacci = &Deck{"Fibonacci", []string{"0", "1", "2", "3", "5", "8", "13", "21", "34", "55", "89", "?", "☕"}}

// TShirtSizes uses a number of shirt sizes for estimates.
var TShirtSizes = &Deck{"T-Shirt Sizes", []string{"XS", "S", "M", "L", "XL", "XXL", "?", "☕"}}

var Hours = &Deck{"Hours", []string{"0", ".5", "1", "2", "4", "8", "12", "16", "20", "24", "?", "☕"}}

// AllDecks contains a mapping of deck names to decks
// type Decks map[string]*Deck

var AllDecks = map[string]*Deck{
	ModifiedFibonacci.Name: ModifiedFibonacci,
	Fibonacci.Name:         Fibonacci,
	TShirtSizes.Name:       TShirtSizes,
	Hours.Name:             Hours,
}

// GetCard returns the card for the specified index.
func (d *Deck) GetCard(i int) (string, error) {
	if i < 0 || i >= len(d.Cards) {
		return "", ErrCardNotFound
	}

	return d.Cards[i], nil
}

func AppendFromJSON(filename string) (err error) {
	newDecks := make(map[string]*Deck)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(buf, &newDecks)
	if err != nil {
		return
	}
	for name, deck := range AllDecks {
		newDecks[name] = deck
	}
	AllDecks = newDecks
	return
}
