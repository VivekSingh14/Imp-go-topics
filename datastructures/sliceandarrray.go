package datastructures

import "fmt"

type Contract1 interface {
	Display()
}

type Shape struct {
	Radius int
}

func (s *Shape) Display() {
	fmt.Println("radius of circle is: ", s.Radius)
}
