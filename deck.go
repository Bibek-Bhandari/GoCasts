package main

import "fmt"
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Ace","Hearts", "clubs"}
	cardValues := []string{"one", "two", "three"}

	for _,cardSuit := range cardSuits{
		for _, cardValue := range cardValues{
			cards = append(cards, cardSuit+ "of"+ cardValue)
		}

	}
	return cards
}
//multiple returns
func deal(d deck,handsize int)(deck,deck){
	return d[:handsize],d[handsize:]
}
//receiver
func (d deck) print() {
	for _,card := range d{
		fmt.Println(card)
	}
}