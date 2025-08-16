package main

import "fmt"

type Point struct {
	id int
	x  float64
	y  float64
}

// Is made changes on a copy
func (p Point) moveSimulation() {
	p.x = 10
	p.y = 20

	fmt.Printf("Point %v moved to x: %v y: %v \n", p.id, p.x, p.y)
}

func (p *Point) move() {
	p.x = 10
	p.y = 20

	fmt.Printf("Point %v moved to x: %v y: %v \n", p.id, p.x, p.y)
}

// Very common
func NewPoint() *Point {
	return &Point{
		id: 0, // Could be a random generator
		x:  0,
		y:  0,
	}
}

func main() {
}
