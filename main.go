package main

func main() {
	cards := deck{"bibek", newcard()}
	cards = append(cards, "bibek")
	cards.print()

}

func newcard() string {
	return "club"
}