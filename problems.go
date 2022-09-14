package main

import (
	"fmt"
	"math"
)

// /////////SQUARE////////////////
type Square struct {
	a int
}

func (v Square) Area() int {
	return (v.a * v.a)
}
func (v Square) Parametr() int {
	return (v.a * 4)
}

//////////TRIANGLE//////////////

type Triangle struct {
	a, b, c float64
}

func (v Triangle) Area() float64 {
	return (0.25 * (math.Sqrt(4*v.a*v.a*v.b*v.b - (math.Pow((v.a*v.a + v.b*v.b - v.c*v.c), 2)))))
}

func (v Triangle) Parametr() float64 {
	return (v.a + v.b + v.c)
}

// ////////RECTENGLE///////////
type Rectengle struct {
	a, b int
}

func (v Rectengle) Area() int {
	return (v.a * v.b)
}

func (v Rectengle) Parametr() int {
	return (2 * (v.a + v.b))
}

// ///////CIRCLE//////////////
type Circle struct {
	a float64
}

func (v Circle) Area() float64 {

	return (math.Pi * v.a * v.a)
}

func (v Circle) Parametr() float64 {
	return (math.Pi * 2 * v.a)
}

func main() {

	v := Circle{5}

	fmt.Println("Area =",v.Area())
	fmt.Println("Parametr =",v.Parametr())

}
