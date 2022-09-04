package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	score := a * b

	if score%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
