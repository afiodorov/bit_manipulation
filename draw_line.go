package main

func mask(start, end int) byte {
	var mask byte
	for i := start; i < end; i++ {
		mask |= (1 << (7 - i))
	}

	return mask
}

func drawLine(screen []byte, width, x1, x2, y int) error {
	numBits := 8
	w := width / numBits
	row := screen[y*w : y*w+w]

	x2++ // because working with half open intervals is nicer: mask [start, end)
	for start := x1; start <= x2; start = ((start / numBits) + 1) * numBits {
		end := ((start / numBits) + 1) * numBits
		if end > x2 {
			end = x2
		}

		length := end - start
		row[start/numBits] |= mask(start%numBits, start%numBits+length)
	}

	return nil
}
