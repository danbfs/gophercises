package deck

import "fmt"

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output: 
	// Ace of Hearts
	// Jack of Clubs
	// Queen of Diamonds
	// Two of Spades
	// Joker
}