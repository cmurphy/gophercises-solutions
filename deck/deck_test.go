package deck

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		d := New()
		if d[0] != (Card{Spade, Ace}) {
			t.Errorf("wrong first card: %v", d[0])
		}
		if d[51] != (Card{Heart, King}) {
			t.Errorf("wrong last card: %v", d[51])
		}
	})
}

func TestShuffle(t *testing.T) {
	randShuffle = rand.New(rand.NewSource(0))
	t.Run("shuffle", func(t *testing.T) {
		res := New(Shuffle())
		if res[0] != (Card{Diamond, Seven}) {
			t.Errorf("expected Seven of Diamonds in position 0, found: %s", res[0])
		}
		if res[51] != (Card{Heart, Jack}) {
			t.Errorf("expected Jack of Hearts in position 51, found: %s", res[51])
		}
	})
}

func TestDefaultSort(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		def := New()
		res := New(Shuffle(), DefaultSort())
		if !reflect.DeepEqual(def, res) {
			t.Errorf("not sorted")
		}
	})
}

func TestCustomSort(t *testing.T) {
	t.Run("custom sort", func(t *testing.T) {
		res := New(Shuffle(), Sort(func(d []Card) func(i, j int) bool { return func(i, j int) bool { return d[i].Rank < d[j].Rank } }))
		for i := 0; i < 4; i++ {
			if int(res[i].Rank) != 0 {
				t.Errorf("not sorted")
			}
		}
	})
}

func TestAddJokers(t *testing.T) {
	t.Run("add jokers", func(t *testing.T) {
		extras := 5
		res := New(AddJokers(extras))
		for i := 52; i < 52+extras; i++ {
			if res[i].Suit != Joker {
				t.Errorf("missing joker")
			}
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter spades", func(t *testing.T) {
		res := New(Filter(func(c Card) bool { return c.Suit == Spade }))
		for _, c := range res {
			if c.Suit == Spade {
				t.Errorf("found a spade")
			}
		}
	})
	t.Run("filter twos", func(t *testing.T) {
		res := New(Filter(func(c Card) bool { return c.Rank == Two }))
		for _, c := range res {
			if c.Rank == Two {
				t.Errorf("found a two")
			}
		}
	})
}

func TestCompose(t *testing.T) {
	t.Run("double deck", func(t *testing.T) {
		res := New(Compose(2))
		if len(res) != 52*2 {
			t.Errorf("wrong size deck")
		}
		if res[52] != (Card{Suit: Spade, Rank: Ace}) {
			t.Errorf("wrong card in slot")
		}
	})
}

func TestString(t *testing.T) {
	testCases := []struct {
		input Card
		want  string
	}{
		{input: Card{Spade, Ace}, want: "Ace of Spades"},
		{input: Card{Spade, Two}, want: "Two of Spades"},
		{input: Card{Spade, Seven}, want: "Seven of Spades"},
		{input: Card{Spade, Ten}, want: "Ten of Spades"},
		{input: Card{Spade, Queen}, want: "Queen of Spades"},
		{input: Card{Diamond, Ace}, want: "Ace of Diamonds"},
		{input: Card{Diamond, Three}, want: "Three of Diamonds"},
		{input: Card{Diamond, Queen}, want: "Queen of Diamonds"},
		{input: Card{Club, Ace}, want: "Ace of Clubs"},
		{input: Card{Club, Four}, want: "Four of Clubs"},
		{input: Card{Club, Five}, want: "Five of Clubs"},
		{input: Card{Club, Nine}, want: "Nine of Clubs"},
		{input: Card{Club, Jack}, want: "Jack of Clubs"},
		{input: Card{Club, King}, want: "King of Clubs"},
		{input: Card{Heart, Ace}, want: "Ace of Hearts"},
		{input: Card{Heart, Six}, want: "Six of Hearts"},
		{input: Card{Heart, Eight}, want: "Eight of Hearts"},
		{input: Card{Heart, Nine}, want: "Nine of Hearts"},
		{input: Card{Heart, Ten}, want: "Ten of Hearts"},
		{input: Card{Heart, Queen}, want: "Queen of Hearts"},
		{input: Card{Suit: Joker}, want: "Joker"},
	}
	for _, tc := range testCases {
		res := tc.input.String()
		if res != tc.want {
			t.Errorf("unexpected value, got: %v, want: %v", res, tc.want)
		}
	}
}
