package main

func createMask() int16 {
	return 1 << 16 / 3
}

func pairWiseSwap(n int16) int16 {
	m := createMask()
	evens := n & m
	odds := n & (^m)

	return evens<<1 | int16(uint16(odds)>>1)
}
