package main

import "fmt"

func main() {
	type (
		UtcTime string
		JstTime string
	)
	var t1 UtcTime = "00:00:00"
	var t2 JstTime = "09:00:00"

	fmt.Println("t1 = ", t1)
	fmt.Println("t2 = ", t2)

	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)

	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
}
