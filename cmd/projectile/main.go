package main

import (
	"fmt"
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
		tuple.Scale(tuple.Normalize(tuple.NewVector(1, 1, 0)), 5),
	)

	newPos := tick(env, p0)
	fmt.Println(newPos)
	// Keep ticking until hitting the ground.
	for newPos.Position.Y >= 0 {
		newPos = tick(env, newPos)
		fmt.Println(newPos)
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
