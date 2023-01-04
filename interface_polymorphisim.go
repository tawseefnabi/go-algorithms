package main

import "fmt"

type Geometry interface {
	Edges() int
}

type Pentagon struct{}
type Hexagon struct{}
type Octagon struct{}
type Decagon struct{}

func (p Pentagon) Edges() int { return 5 }
func (h Hexagon) Edges() int  { return 6 }
func (o Octagon) Edges() int  { return 8 }
func (d Decagon) Edges() int  { return 10 }

// Parameter calculate parameter of object
func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

func main() {
	// p := new(Pentagon)
	var p Pentagon
	fmt.Println(p.Edges())
	fmt.Println(Parameter(p, 9))

	var h Hexagon
	fmt.Println(h.Edges())
	fmt.Println(Parameter(h, 9))
}
