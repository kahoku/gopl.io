// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 106.

// Embed demonstrates basic struct embedding.
package main

import "fmt"

type Point struct{ X, Y int }

func (p *Point) draw() {
	fmt.Printf("Point.draw:X=%v,Y=%v\n", p.X, p.Y)
}

type Circle struct {
	Point
	Radius int
}

func (c *Circle) draw() {
	fmt.Printf("Circle.draw:X=%v,Y=%v,Radius=%v\n", c.X, c.Y, c.Radius)
}

type Wheel struct {
	Circle
	Spokes int
}
type Triangle struct {
	P1 Point
	P2 Point
	P3 Point
}

func main() {
	var w Wheel
	//!+
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
	//!-

	w.draw()

	tri := Triangle{P1: Point{1, 1}, P2: Point{1, 2}, P3: Point{2, 1}}
	fmt.Printf("%#v\n", tri)
	//outPut main.Triangle{P1:main.Point{X:1, Y:1}, P2:main.Point{X:1, Y:2}, P3:main.Point{X:2, Y:1}}
}
