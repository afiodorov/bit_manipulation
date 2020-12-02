package main

import (
	"math"
	"testing"
)

func TestNextLargest(t *testing.T) {
	for i := uint16(0); i < math.MaxUint16; i++ {
		res1, ok1 := nextLargestNaive(i)
		res2, ok2 := nextLargestMask(i)

		if res1 != res2 || ok1 != ok2 {
			t.Fatalf("fail for input %v: (%v, %v) vs (%v, %v)\n", i, res1, ok1, res2, ok2)
		}
	}
}

func TestNextSmallestNaive(t *testing.T) {
	if _, ok := nextSmallestNaive(0); ok {
		t.Fail()
	}

	if _, ok := nextSmallestNaive(1); ok {
		t.Fail()
	}

	if r, ok := nextSmallestNaive(2); !ok || r != 1 {
		t.Fatalf("Got (%v, %v)", r, ok)
	}

	if r, ok := nextSmallestNaive(5); !ok || r != 3 {
		t.Fatalf("Got (%v, %v)", r, ok)
	}
}

func TestNextSmallest(t *testing.T) {
	for i := uint16(0); i < math.MaxUint16; i++ {
		res1, ok1 := nextSmallestNaive(i)
		res2, ok2 := nextSmallestMask(i)

		if res1 != res2 || ok1 != ok2 {
			t.Fatalf("fail for input %v: (%v, %v) vs (%v, %v)\n", i, res1, ok1, res2, ok2)
		}
	}
}
