//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

func (c *Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

type Option func(deck []Card) []Card

func New(options ...Option) []Card {
	cards := make([]Card, 52)
	for s := Spade; s <= Heart; s++ {
		for r := Ace; r <= King; r++ {
			cards[int(s)*(int(King)+1)+int(r)] = Card{s, r}
		}
	}
	for _, o := range options {
		var err error
		cards = o(cards)
		if err != nil {
			return cards
		}
	}
	return cards
}

var randShuffle = rand.New(rand.NewSource(time.Now().UnixNano()))

func Shuffle() Option {
	return func(cards []Card) []Card {
		//rand.Seed(time.Now().UnixNano())
		randShuffle.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].Suit > cards[j].Suit {
			return false
		}
		if cards[i].Suit < cards[j].Suit {
			return true
		}
		return cards[i].Rank < cards[j].Rank
	}
}

func Sort(less func(cards []Card) func(i, j int) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort() Option {
	return Sort(Less)
}

func AddJokers(n int) Option {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) Option {
	return func(cards []Card) []Card {
		d := []Card{}
		for _, c := range cards {
			if !f(c) {
				d = append(d, c)
			}
		}
		return d
	}
}

func Compose(n int) Option {
	return func(cards []Card) []Card {
		d := []Card{}
		for i := 0; i < n; i++ {
			d = append(d, cards...)
		}
		return d
	}
}
