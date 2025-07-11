package encounter

import (
	"math/rand"
)

type Deck struct {
	cards    []int
	position int
}

func (deck *Deck) Suffle() {
	// Shuffle using Fisher-Yates
	for i := len(deck.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
	deck.position = 0
}

func (deck *Deck) Draw() int {
	var card int
	if deck.position >= 54 {
		deck.Suffle()
	}
	card = deck.cards[deck.position]
	deck.position++

	return card
}

func NewDeck() Deck {
	// Create array with numbers 0-53
	arr := make([]int, 54)
	for i := 0; i < 54; i++ {
		arr[i] = i
	}

	// Shuffle using Fisher-Yates
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	deck := Deck{
		cards:    arr,
		position: 0,
	}
	return deck
}
