package main

import (
	"fmt"
)

type Student struct {
	name    string
	math    float64
	english float64
}

func (s Student) avg(math, english float64) float64 {
	return (s.math + s.english) / 2
}

func main() {
	var s Student
	s.name = "sato"
	s.math = 80
	s.english = 70

	fmt.Println(s)
	fmt.Println(s.avg(s.math, s.english))
}
