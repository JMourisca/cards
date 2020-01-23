package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 52
	if len(d) != expected {
		t.Errorf("Expected %d got %d", expected, len(d))
	}
}