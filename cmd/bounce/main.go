package main

import (
	"fmt"
	"log"
	"os"
	"rt/canvas"
	"rt/tuple"
	"time"
)

const (
	CanvasWidth  = 1024
	CanvasHeight = 768
)

func main() {
	c := canvas.NewCanvas(CanvasWidth, CanvasHeight, 255)
	file, err := os.Create("bounce.ppm")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	ticker := time.NewTicker(time.Millisecond * 5)
	defer ticker.Stop()

	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 3)
		done <- struct{}{}
	}()

	x, y := 0, 0
	dx, dy := 5, 5
	color := tuple.NewColor(1, 0, 0)

	completed := false
	for !completed {
		select {
		case <-done:
			completed = true
			ticker.Stop()
		case <-ticker.C:
			c.WritePixel(x, y, color)
			if x+dx > CanvasWidth || x+dx < 0 {
				dx = -dx
				color = tuple.NewColor(0, 1, 0)
			}
			if y+dy > CanvasHeight || y+dy < 0 {
				dy = -dy
				color = tuple.NewColor(0, 0, 1)
			}
			x += dx
			y += dy
		}
	}

	fmt.Println("Done")
	file.WriteString(c.ToPPM())
}
