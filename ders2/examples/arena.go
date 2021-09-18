package examples

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Square struct {
	Width  float64
	Height float64
}

func (s *Square) Area() float64 {
	return s.Width * s.Height
}

type Sizer interface {
	Area() float64
}

func Eq(s1, s2 Sizer) {
	if s1.Area() < s2.Area() {
		fmt.Println("s1 less")
	} else {
		fmt.Println("s2 less")
	}
}
func Main_arena() {
	c := Circle{Radius: 10}
	fmt.Println(c)

	s := Square{Width: 10, Height: 8}
	fmt.Println(s)

	/* if c.Area() < s.Area() {
		fmt.Println("Circle less")
	} else {
		fmt.Println(c.Area())
	} */

	Eq(&c, &s)

}
