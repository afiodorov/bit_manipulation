package main

import (
	"strings"
)

func binary2string(x float64) string {
	errString := "ERROR"

	if x >= 1 || x <= 0 {
		return errString
	}

	var sb strings.Builder
	if _, err := sb.WriteRune('.'); err != nil {
		return errString
	}

	for x > 0 {
		if sb.Len() >= 32 {
			return errString
		}

		r := x * 2
		if r >= 1 {
			if _, err := sb.WriteRune('1'); err != nil {
				return errString
			}

			x = r - 1
		} else {
			if _, err := sb.WriteRune('0'); err != nil {
				return errString
			}

			x = r
		}
	}

	return sb.String()
}
