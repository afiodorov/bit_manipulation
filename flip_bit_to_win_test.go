package main

import "testing"

func TestFlipBitToWin(t *testing.T) {
	if flipBitToWin(1775) != 8 {
		t.Fail()
	}
}
