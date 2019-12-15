package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Output:
	// Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("Expected %d cards, got %d", 13*4, len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Errorf("Expected %s as first card. Got %s", exp, cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Errorf("Expected %s as first card. Got %s", exp, cards[0])
	}
}

func TestJokers(t *testing.T) {
	desired := 3
	cards := New(Jokers(desired))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != desired {
		t.Errorf("Expected %d jokers. Got %d", desired, count)
	}
}
