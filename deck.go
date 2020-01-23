package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Createa a new type of 'deck'
// which is a slice of strings
type deck []string // Kind of inheriting the type string "class deck extends []string"

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♠", "♦", "♥", "♣"}

	for _, suit := range cardSuits {
		for i := 0; i < 13; i++ {
			value := strconv.Itoa(i + 1)
			switch value {
			case "1":
				value = "A"
			case "11":
				value = "J"
			case "12":
				value = "Q"
			case "13":
				value = "K"
			}

			cardValue := fmt.Sprintf("%s%s", value, suit)
			cards = append(cards, cardValue)
		}
	}

	return cards
}

func (d deck) count() {
	fmt.Println(len(d))
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}