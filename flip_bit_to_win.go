package main

func getBit(x int64, i int) bool {
	return (x & (1 << i)) != 0
}

func flipBitToWin(x int64) int64 {
	var pos []int

	for i := 0; i < 64; i++ {
		if !getBit(x, i) {
			pos = append(pos, i)
		}
	}

	max := -1

	for i, p := range pos {
		var lenLeft int
		if i >= 1 {
			lenLeft = (p - pos[i-1]) - 1
		} else {
			lenLeft = p
		}

		var lenRight int
		if i+1 < len(pos) {
			lenRight = pos[i+1] - p - 1
		} else {
			lenRight = 0
		}

		if lenLeft+lenRight+1 > max {
			max = lenLeft + lenRight + 1
		}
	}

	return int64(max)
}
