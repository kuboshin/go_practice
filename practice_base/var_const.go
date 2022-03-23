package main

import "fmt"

func main() {
	var a1 int = 123
	var a2 = 123
	a3 := 123

	var (
		a4 int = 123
		a5 int = 456
	)

	a1 = 456
	var name, age = "kubo", 30

	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("a3 = ", a3)
	fmt.Println("a4 = ", a4)
	fmt.Println("a5 = ", a5)
	fmt.Println("name = ", name)
	fmt.Println("name = ", age)

	const c1 = 123
	const (
		c2 = 123
		c3 = 456
	)
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	fmt.Println("c3 = ", c3)

}
