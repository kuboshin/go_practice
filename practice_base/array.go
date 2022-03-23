package main

import "fmt"

func main() {
	a1 := [3]string{}
	a1[0] = "Red"
	a1[1] = "Green"
	a1[2] = "Blue"
	fmt.Println(a1[0], a1[1], a1[2])

	a2 := [...]string{"red", "green", "blue"}
	fmt.Println(a2[0], a2[1], a2[2])

	a3 := []string{}
	a3 = append(a3, "Red")
	a3 = append(a3, "Green")
	a3 = append(a3, "Blue")
	fmt.Println(a3[0], a3[1], a3[2])

	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}
}
