package main

import (
	"fmt"
)

func main() {
	fmt.Println("Good Morining")
	fmt.Println("Good Afternoon")
	fmt.Println("Good Evening")

	a := [...]string{"sato", "suzuki", "yamada"}

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}
