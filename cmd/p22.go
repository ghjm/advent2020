package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"strings"
)

func sumOfDeck(deck []int) int64 {
	var sum int64
	for i := 0; i < len(deck); i++ {
		sum += int64(deck[len(deck)-i-1]) * int64(i + 1)
	}
	return sum
}

//var gameno = 0

func playCombat(p1Deck []int, p2Deck []int, recursive bool) (int, int64) {
	//gameno++
	//mygame := gameno
	//fmt.Printf("=== Game %d ===\n", mygame)
	previousPositions := make(map[string]struct{})
	var winner int
	round := 0
	for {
		round++
		//fmt.Printf("-- Round %d (Game %d) --\n", round, mygame)
		//fmt.Printf("Player 1: %v\n", p1Deck)
		//fmt.Printf("Player 2: %v\n", p2Deck)
		curPosition := fmt.Sprintf("%v %v", p1Deck, p2Deck)
		_, ok := previousPositions[curPosition]
		if ok {
			//fmt.Printf("Player 1 wins due to repeat of previous position\n")
			return 1, sumOfDeck(p1Deck)
		}
		previousPositions[curPosition] = struct{}{}
		p1Card := p1Deck[0]
		p1Deck = p1Deck[1:]
		p2Card := p2Deck[0]
		p2Deck = p2Deck[1:]
		//fmt.Printf("Player 1 plays: %d\n", p1Card)
		//fmt.Printf("Player 2 plays: %d\n", p2Card)
		if recursive && len(p1Deck) >= p1Card && len(p2Deck) >= p2Card {
			//fmt.Printf("Playing a sub-game to determine the winner...\n")
			p1SubDeck := make([]int, p1Card)
			for i := 0; i < p1Card; i++ {
				p1SubDeck[i] = p1Deck[i]
			}
			p2SubDeck := make([]int, p2Card)
			for i := 0; i < p2Card; i++ {
				p2SubDeck[i] = p2Deck[i]
			}
			winner, _ = playCombat(p1SubDeck, p2SubDeck, true)
			//fmt.Printf("Back to game %d...\n", mygame)
		} else if p1Card > p2Card {
			winner = 1
		} else {
			winner = 2
		}
		//fmt.Printf("Player %d wins!\n", winner)
		if winner == 1 {
			p1Deck = append(p1Deck, p1Card, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card, p1Card)
		}
		if len(p1Deck) == 0 || len(p2Deck) == 0 {
			break
		}
	}
	var sum int64
	if len(p1Deck) == 0 {
		winner = 2
		sum = sumOfDeck(p2Deck)
	} else {
		winner = 1
		sum = sumOfDeck(p1Deck)
	}
	return winner, sum
}

func main() {
	cards := make([][]int, 0)
	var deck []int
	err := utils.OpenAndReadAll("input22.txt", func(s string) error {
		if s == "" {
			return nil
		} else if strings.HasPrefix(s, "Player") {
			if deck != nil {
				cards = append(cards, deck)
			}
			deck = make([]int, 0)
		} else {
			card := utils.AtoiOrPanic(s)
			deck = append(deck, card)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	if deck != nil {
		cards = append(cards, deck)
	}
	fmt.Printf("Part One\n")
	if len(cards) != 2 {
		panic("wrong number of decks")
	}
	p1Deck := make([]int, len(cards[0]))
	for i := range cards[0] {
		p1Deck[i] = cards[0][i]
	}
	p2Deck := make([]int, len(cards[1]))
	for i := range cards[1] {
		p2Deck[i] = cards[1][i]
	}
	var sum int64
	//_, sum := playCombat(p1Deck, p2Deck, false)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Part Two\n")
	p1Deck = make([]int, len(cards[0]))
	for i := range cards[0] {
		p1Deck[i] = cards[0][i]
	}
	p2Deck = make([]int, len(cards[1]))
	for i := range cards[1] {
		p2Deck[i] = cards[1][i]
	}
	_, sum = playCombat(p1Deck, p2Deck, true)
	fmt.Printf("Sum: %d\n", sum)
}

