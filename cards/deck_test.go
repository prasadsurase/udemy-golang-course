package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace pf Hearts. Got %v", d[0])
	}

	if d[15] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs. Got %v", d[15])
	}
}
