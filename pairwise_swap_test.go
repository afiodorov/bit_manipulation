package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestCreateMask(t *testing.T) {
	if strconv.FormatInt(int64(createMask()), 2) != "101010101010101" {
		t.Fail()
	}
}

func TestPairWiseSwapSimple(t *testing.T) {
	if pairWiseSwap(1) != 2 {
		t.Fail()
	}

	if pairWiseSwap(3) != 3 {
		t.Fail()
	}

	if pairWiseSwap(2) != 1 {
		t.Fail()
	}

	if pairWiseSwap(10) != 5 { // 1010 -> 0101
		t.Fail()
	}
}

func int2string(n int16) string {
	if n >= 0 {
		return fmt.Sprintf("%016b", n)
	}
	return fmt.Sprintf("%016b", 1<<16+int64(n))
}

func swap(in string) string {
	var sb strings.Builder
	runes := []rune(in)

	for i := 0; i < len(runes); i += 2 {
		sb.WriteRune(runes[i+1])
		sb.WriteRune(runes[i])
	}

	return sb.String()
}

func TestPairWiseSwapExhaustive(t *testing.T) {
	for n := math.MinInt16; n <= math.MaxInt16; n++ {
		r := pairWiseSwap(int16(n))
		input := int2string(int16(n))
		swapped := int2string(r)
		if swap(input) != swapped {
			t.Fatalf("Failed for %v. Got: %v, Input: %v. Expected: %v", n, r, swap(input), swapped)
		}
	}
}
