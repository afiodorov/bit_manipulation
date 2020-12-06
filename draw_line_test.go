package main

import (
	"fmt"
	"strings"
	"testing"
)

func byte2string(b byte) string {
	return fmt.Sprintf("%08b", b)
}

func renderScreen(screen []byte, width int) string {
	var sb strings.Builder

	w := width / 8

	for i := 0; i < len(screen); i += w {
		row := screen[i : i+w]
		for _, b := range row {
			sb.WriteString(byte2string(b))
		}

		sb.WriteRune('\n')
	}

	return strings.TrimSpace(sb.String())
}

func TestRenderScreen(t *testing.T) {
	for _, tc := range []struct {
		screen   []byte
		width    int
		expected string
	}{
		{screen: []byte{0xFF, 0x0F, 0x10},
			width: 8,
			expected: strings.TrimSpace(`
11111111
00001111
00010000`)},
		{screen: []byte{0xFF, 0xFF, 0x00, 0x00, 0x10, 0x10},
			width: 16,
			expected: strings.TrimSpace(`
1111111111111111
0000000000000000
0001000000010000`)},
	} {
		actual := renderScreen(tc.screen, tc.width)
		if tc.expected != actual {
			t.Fatalf("Got:\n%v\nExpected:\n%v", actual, tc.expected)
		}
	}
}

func TestDrawLine(t *testing.T) {
	for _, tc := range []struct {
		screen   []byte
		width    int
		x1       int
		x2       int
		y        int
		expected string
	}{
		{screen: []byte{0x00, 0x00, 0x00},
			width: 8,
			x1:    2,
			x2:    4,
			y:     1,
			expected: strings.TrimSpace(`
00000000
00111000
00000000`)},
	} {
		drawLine(tc.screen, tc.width, tc.x1, tc.x2, tc.y)
		actual := renderScreen(tc.screen, tc.width)
		if tc.expected != actual {
			t.Fatalf("Got:\n%v\nExpected:\n%v", actual, tc.expected)
		}
	}
}
