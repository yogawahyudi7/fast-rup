package main

import (
	"fmt"
	"math"
)

type calculate interface {
	area() float64
	perimeter() float64
}

type circle struct {
	diameter float64
}

func (l circle) radius() float64 {
	return l.diameter / 2
}

func (l circle) area() float64 {
	return math.Pi * math.Pow(l.radius(), 2)
}

func (l circle) perimeter() float64 {
	return math.Pi * l.diameter
}

type square struct {
	side float64
}

func (p square) area() float64 {
	return math.Pow(p.side, 2)
}

func (p square) perimeter() float64 {
	return p.side * 4
}

func main() {
	var flatBuilding calculate

	flatBuilding = square{10.0}
	fmt.Println("===== square")
	fmt.Println("area      :", flatBuilding.area())
	fmt.Println("perimeter  :", flatBuilding.perimeter())

	flatBuilding = circle{14.0}
	fmt.Println("===== circle")
	fmt.Println("area      :", flatBuilding.area())
	fmt.Println("perimeter  :", flatBuilding.perimeter())
	fmt.Println("radius :", flatBuilding.(circle).radius())
}
