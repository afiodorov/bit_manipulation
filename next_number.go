package main

import (
	"math"
)

func updateBit(x uint16, i int, bitIs1 bool) uint16 {
	var value uint16
	if bitIs1 {
		value = 1
	}

	mask := ^(uint16(1) << i)

	return (x & mask) | (value << i)
}

func numOneBits(x uint16) int16 {
	var count int16

	for x > 0 {
		count += int16(x & uint16(1))
		x >>= 1
	}

	return count
}

func safeIncr(num uint16) (uint16, bool) {
	inc := num + 1

	if num == math.MaxUint16 {
		return inc, false
	}

	return inc, true
}

func safeDecr(num uint16) (uint16, bool) {
	dec := num - 1

	if num == 0 {
		return num, false
	}

	return dec, true
}

func nextLargestNaive(num uint16) (res uint16, ok bool) {
	numBits := numOneBits(num)

	for res, ok = safeIncr(num); (numOneBits(res) != numBits) && ok; res, ok = safeIncr(res) {
	}

	return res, ok
}

func nextLargestMask(num uint16) (res uint16, ok bool) {
	x := num

	pos := 0
	numOnes := 0
	found := false

	for x > 0 {
		if x%4 == 1 {
			found = true
			break
		}

		numOnes += int(x & uint16(1))

		x >>= 1
		pos++
	}

	if !found || pos == 16-1 {
		return res, false
	}

	res = updateBit(num, pos, false)
	res = updateBit(res, pos+1, true)

	for i := 0; i < pos; i++ {
		res = updateBit(res, i, i < numOnes)
	}

	return res, true
}

func nextSmallestNaive(num uint16) (res uint16, ok bool) {
	numBits := numOneBits(num)

	for res, ok = safeDecr(num); (numOneBits(res) != numBits) && ok; res, ok = safeDecr(res) {
	}

	return res, ok
}

func nextSmallestMask(num uint16) (res uint16, ok bool) {
	x := num

	pos := 0
	numZeros := 0
	found := false

	for x > 0 {
		if x%4 == 2 {
			found = true
			break
		}

		if x%2 == 0 {
			numZeros++
		}

		x >>= 1
		pos++
	}

	if !found || pos == 16-1 {
		return res, false
	}

	res = updateBit(num, pos, true)
	res = updateBit(res, pos+1, false)

	for i := 0; i < pos; i++ {
		res = updateBit(res, i, i >= numZeros)
	}

	return res, true
}
