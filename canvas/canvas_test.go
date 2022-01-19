package canvas

import (
	"bufio"
	"rt/tuple"
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20, 255)
	if c.width != 10 {
		t.Fatalf("Wrong width: got: %d; want: %d", c.width, 10)
	}
	if c.height != 20 {
		t.Fatalf("Wrong height: got: %d; want: %d", c.height, 20)
	}
	blackPixel := tuple.NewColor(0, 0, 0)
	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			if !c.At(x, y).IsEqual(blackPixel) {
				t.Fatalf("Pixel at %d,%d is not black", x, y)
			}
		}
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	c := NewCanvas(10, 20, 255)
	red := tuple.NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	actual := c.At(2, 3)
	if !actual.IsEqual(red) {
		t.Fatalf("want: %v; got: %v", red, actual)
	}
}

func TestCanvas_ToPPM(t *testing.T) {
	c := NewCanvas(5, 3, 255)
	c.WritePixel(0, 0, tuple.NewColor(1.5, 0, 0))
	c.WritePixel(2, 1, tuple.NewColor(0, 0.5, 0))
	c.WritePixel(4, 2, tuple.NewColor(-0.5, 0, 1))

	actual := c.ToPPM()
	expectedHeader := "P3\n5 3\n255\n"
	expectedContent := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	reader := strings.NewReader(actual)
	scanner := bufio.NewScanner(reader)
	header := ""
	for i := 0; i < 3 && scanner.Scan(); i++ {
		header += scanner.Text()
		header += "\n"
	}
	if header != expectedHeader {
		t.Fatalf("want: %s; got: %s", expectedHeader, header)
	}

	content := ""
	for i := 0; i < 3 && scanner.Scan(); i++ {
		content += scanner.Text()
		content += "\n"
	}
	if content != expectedContent {
		t.Fatalf("want: %s; got: %s", expectedContent, content)
	}
}

func TestCanvas_MapToRange(t *testing.T) {
	testCases := []struct {
		name                         string
		inputNumber                  float64
		inputOldRange, inputNewRange Range
		expected                     int64
	}{
		{
			name:          "5, [0,10] -> [10,20]",
			inputNumber:   5,
			inputOldRange: Range{0, 10},
			inputNewRange: Range{10, 20},
			expected:      15,
		},
		{
			name:          "0, [0,10] -> [10,20]",
			inputNumber:   0,
			inputOldRange: Range{0, 10},
			inputNewRange: Range{10, 20},
			expected:      10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MapToRange(tc.inputOldRange, tc.inputNewRange, tc.inputNumber)
			if actual != tc.expected {
				t.Fatalf("got: %d; expected: %d", actual, tc.expected)
			}
		})
	}
}
