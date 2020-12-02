package main

import (
	"math"
	"testing"
)

func TestBinary2string(t *testing.T) {
	if binary2string(math.Pow(2, -30)+math.Pow(2, -15)) != ".000000000000001000000000000001" {
		t.Fail()
	}
}
