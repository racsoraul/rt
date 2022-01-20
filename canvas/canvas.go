package canvas

import (
	"fmt"
	"math"
	"rt/tuple"
	"strconv"
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
	if x <= c.width && y <= c.height {
		c.pixels[x][y] = color
	}
}

// ToPPM Returns a Plain Portable Pixel Map format.
func (c *Canvas) ToPPM() string {
	// Creates header of the PPM.
	header := fmt.Sprintf("P3\n%d %d\n%d\n", c.width, c.height, c.maxColorValue)

	var content, line strings.Builder
	content.WriteString(header)

	// Initializes ranges to map pixel values.
	originalRange := Range{0, 1}
	newRange := Range{0, 255}

	var triplet, r, g, b string

	// Creates body of the image.
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			r = strconv.Itoa(MapToRange(originalRange, newRange, clamp(c.pixels[x][y][0])))
			g = strconv.Itoa(MapToRange(originalRange, newRange, clamp(c.pixels[x][y][1])))
			b = strconv.Itoa(MapToRange(originalRange, newRange, clamp(c.pixels[x][y][2])))
			triplet = fmt.Sprintf("%s %s %s", r, g, b)

			if line.Len() == 70 {
				content.WriteString(line.String())
				content.WriteByte('\n')
				line.Reset()
			}
			if line.Len()+len(triplet)+1 <= 70 {
				if line.Len() != 0 {
					line.WriteByte(' ')
				}
				line.WriteString(triplet)
				continue
			}
			if line.Len()+len(r)+1 > 70 {
				content.WriteString(line.String())
				content.WriteByte('\n')
				line.Reset()
				line.WriteString(triplet)
				continue
			}
			line.WriteByte(' ')
			line.WriteString(r)

			if line.Len()+len(g)+1 > 70 {
				content.WriteString(line.String())
				content.WriteByte('\n')
				line.Reset()
				line.WriteString(fmt.Sprintf("%s %s", g, b))
				continue
			}
			line.WriteByte(' ')
			line.WriteString(g)

			if line.Len()+len(b)+1 > 70 {
				content.WriteString(line.String())
				content.WriteByte('\n')
				line.Reset()
				line.WriteString(b)
				continue
			}
			line.WriteByte(' ')
			line.WriteString(b)
		}
		content.WriteString(line.String())
		content.WriteByte('\n')
		line.Reset()
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
func MapToRange(old, new Range, n float64) int {
	return int(math.Round(new.Min + (n-old.Min)*(new.Max-new.Min)/(old.Max-old.Min)))
}
