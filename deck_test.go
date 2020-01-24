package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 52
	if len(d) != expected {
		t.Errorf("Expected %d got %d", expected, len(d))
	}

	expectedFirstCard := "A♠"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %s got %s", expectedFirstCard, d[0])
	}

	expectedLastCard := "K♣"
	if d[51] != expectedLastCard {
		t.Errorf("Expected last card to be %s got %s", expectedLastCard, d[51])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")
	deck := newDeck()
	_ = deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	expected := 52
	if len(loadedDeck) != expected {
		t.Errorf("Expected %d got %d", expected, len(loadedDeck))
	}
	_ = os.Remove("_decktesting")
}