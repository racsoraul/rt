package canvas

import (
	"fmt"
	"math"
	"rt/tuple"
	"strings"
)

type Canvas struct {
	width         int
	height        int
	maxColorValue int
	pixels        [][]tuple.Tuple
}

// NewCanvas Creates a black new canvas.
func NewCanvas(width, height, maxColorValue int) *Canvas {
	c := &Canvas{width, height, maxColorValue, make([][]tuple.Tuple, width)}
	for i := 0; i < width; i++ {
		c.pixels[i] = make([]tuple.Tuple, height)
	}
	return c
}

// At Returns the pixel at the given coordinates.
func (c *Canvas) At(x, y int) tuple.Tuple {
	return c.pixels[x][y]
}

// WritePixel Writes a pixel at the given coordinates.
func (c *Canvas) WritePixel(x, y int, color tuple.Tuple) {
	c.pixels[x][y] = color
}

// ToPPM Returns a Plain Portable Pixel Map format.
func (c *Canvas) ToPPM() string {
	// Creates header of the PPM.
	header := fmt.Sprintf("P3\n%d %d\n%d\n", c.width, c.height, c.maxColorValue)

	var content strings.Builder
	content.WriteString(header)

	// Initializes ranges to map pixel values.
	originalRange := Range{0, 1}
	newRange := Range{0, 255}

	// Creates body of the PPM.
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			fmt.Fprintf(
				&content,
				"%d %d %d",
				MapToRange(originalRange, newRange, clamp(c.pixels[x][y][0])),
				MapToRange(originalRange, newRange, clamp(c.pixels[x][y][1])),
				MapToRange(originalRange, newRange, clamp(c.pixels[x][y][2])),
			)
			if x != c.width-1 {
				content.WriteString(" ")
			}
		}
		content.WriteString("\n")
	}
	return content.String()
}

// clamp Clamps value to range 0 - 1. If lower than zero, it returns 0.
// If greater than one, it returns 1.
func clamp(value float64) float64 {
	if tuple.EqualFloat(value, 0) || value < 0 {
		return 0
	}
	if tuple.EqualFloat(value, 1) || value > 1 {
		return 1
	}
	return value
}

type Range struct {
	Min, Max float64
}

// MapToRange Maps n number that falls between old range and returns
// a new number that falls between the new range.
func MapToRange(old, new Range, n float64) int64 {
	return int64(math.Round(new.Min + (n-old.Min)*(new.Max-new.Min)/(old.Max-old.Min)))
}
