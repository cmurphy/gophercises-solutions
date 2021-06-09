package main

import (
	"fmt"
	"strings"

	"github.com/cmurphy/gophercises/deck"
)

type Hand []deck.Card

type Stage int8

const (
	NewGame Stage = iota
	PlayerTurn
	DealerTurn
	Done
)

type GameState struct {
	Deck   []deck.Card
	Player Hand
	Dealer Hand
	Stage  Stage
}

func main() {
	game := NewGameState()
	fmt.Printf("Your hand: %s\n", game.Player.String())
	playerScore, _ := score(game.Player)
	fmt.Printf("Your score: %d\n", playerScore)
	fmt.Printf("Dealer hand: %s\n", game.Dealer.DealerString())
	var input string
	// Player turn
	game = game.Next()
	for game.Stage == PlayerTurn {
		fmt.Printf("hit or stand? : ")
		fmt.Scanln(&input)
		switch input {
		case "stand":
			game = game.Next()
		case "hit":
			game = game.Hit()
			fmt.Printf("Your hand: %s\n", game.Player.String())
			playerScore, _ = score(game.Player)
			fmt.Printf("Your score: %d\n", playerScore)
		}
	}
	// Dealer turn
	for game.Stage == DealerTurn {
		dealerScore, soft := score(game.Dealer)
		if dealerScore <= 16 || (soft && dealerScore <= 17) {
			game = game.Hit()
		} else {
			game = game.Next()
		}
	}

	fmt.Printf("Your hand: %s\n", game.Player.String())
	playerScore, _ = score(game.Player)
	fmt.Printf("Your score: %d\n", playerScore)
	fmt.Printf("Dealer hand: %s\n", game.Dealer.String())
	dealerScore, _ := score(game.Dealer)
	fmt.Printf("Dealer score: %d\n", dealerScore)
	switch {
	case playerScore > 21:
		fmt.Println("You're busted!")
	case dealerScore > 21:
		fmt.Println("Dealer's busted!")
	case playerScore > dealerScore:
		fmt.Println("You win!")
	case dealerScore > playerScore:
		fmt.Println("Dealer wins!")
	case playerScore == dealerScore:
		fmt.Println("Draw!")
	}
}

func NewGameState() GameState {
	g := GameState{
		Deck:   deck.New(deck.Shuffle()),
		Dealer: make(Hand, 2),
		Player: make(Hand, 2),
		Stage:  NewGame,
	}
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&g.Player, &g.Dealer} {
			(*hand)[i], g.Deck = deal(g.Deck)
		}
	}
	return g
}

func clone(g GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(g.Deck)),
		Player: make(Hand, len(g.Player)),
		Dealer: make(Hand, len(g.Dealer)),
		Stage:  g.Stage,
	}
	copy(ret.Deck, g.Deck)
	copy(ret.Player, g.Player)
	copy(ret.Dealer, g.Dealer)
	return ret
}

func (g *GameState) CurrentPlayer() *Hand {
	switch g.Stage {
	case PlayerTurn:
		return &g.Player
	case DealerTurn:
		return &g.Dealer
	default:
		panic("no current player")
	}
}

func (g GameState) Next() GameState {
	res := clone(g)
	res.Stage++
	return res
}

func (g GameState) Hit() GameState {
	res := clone(g)
	hand := res.CurrentPlayer()
	var card deck.Card
	card, res.Deck = deal(res.Deck)
	*hand = append(*hand, card)
	return res
}

func deal(d []deck.Card) (deck.Card, []deck.Card) {
	return d[0], d[1:]
}

func (h Hand) String() string {
	res := make([]string, len(h))
	for i, c := range h {
		res[i] = c.String()
	}
	return strings.Join(res, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", ***"
}

func score(h Hand) (int, bool) {
	points := 0
	aces := 0
	soft := false
	for _, c := range h {
		switch {
		case c.Rank == deck.Ace:
			aces++
			continue
		case c.Rank >= deck.Two && c.Rank <= deck.Ten:
			points += int(c.Rank) + 1
		case c.Rank >= deck.Jack && c.Rank <= deck.King:
			points += 10
		}
	}
	for i := 0; i < aces; i++ {
		if points+11 > 21 {
			points += 1
		} else {
			points += 11
			soft = true
		}
	}
	return points, soft
}
