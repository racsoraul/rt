package main

import (
	"fmt"
	"log"
	"os"
	"rt/canvas"
	"rt/tuple"
)

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

func NewProjectile(position, velocity tuple.Tuple) Projectile {
	return Projectile{position, velocity}
}

type Environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func NewEnvironment(gravity, wind tuple.Tuple) Environment {
	return Environment{gravity, wind}
}

func main() {
	env := NewEnvironment(tuple.NewVector(0, -0.1, 0), tuple.NewVector(-0.01, 0, 0))
	p0 := NewProjectile(
		tuple.NewPoint(0, 1, 0),
		tuple.Scale(tuple.Normalize(tuple.NewVector(1, 1, 0)), 10),
	)
	const WIDTH, HEIGHT int = 1024, 260
	c := canvas.NewCanvas(WIDTH, HEIGHT, 255)

	proj := tick(env, p0)
	fmt.Println(proj)
	// Keep ticking until hitting the ground.
	for proj.Position[1] >= 0 {
		proj = tick(env, proj)
		c.WritePixel(int(proj.Position[0]), HEIGHT-int(proj.Position[1]), tuple.NewColor(1, 0, 0))
		fmt.Println(proj)
	}

	file, err := os.Create("projectile.ppm")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString(c.ToPPM())
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(e Environment, p Projectile) Projectile {
	pos, err := tuple.Add(p.Position, p.Velocity)
	if err != nil {
		panic(err)
	}
	velGra, err := tuple.Add(p.Velocity, e.Gravity)
	if err != nil {
		panic(err)
	}
	vel, err := tuple.Add(velGra, e.Wind)
	if err != nil {
		panic(err)
	}
	return NewProjectile(pos, vel)
}
