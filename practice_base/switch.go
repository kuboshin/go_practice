package main

import (
	"fmt"
	"time"
)

// メインの処理を記述していきます
func main() {
	mode := "running"
	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}

	x := 1
	y := 1
	switch {
	case x > y:
		fmt.Println("Big")
	case x < y:
		fmt.Println("Small")
	default:
		fmt.Println("Equal")
	}

	youbi := [...]string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	num := time.Now().Weekday()
	fmt.Println("num:", num)
	dayOfWeek := youbi[num]
	fmt.Println("youbi:", dayOfWeek)
	switch dayOfWeek {
	case "sat":
		fallthrough
	case "sun":
		fmt.Println("Horiday")
	default:
		fmt.Println("Weekday")
	}
}
