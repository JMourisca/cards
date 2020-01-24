// Cards game from udemy course https://www.udemy.com/course/go-the-complete-developers-guide
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}