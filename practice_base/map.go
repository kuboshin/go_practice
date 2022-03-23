package main

import "fmt"

func main() {
	m1 := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(m1["x"])

	m1["z"] = 300
	delete(m1, "x")
	fmt.Println(len(m1))

	// マップに要素が存在するかどうかを調べる
	_, ok := m1["z"]
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not exist")
	}

	// マップをループ処理する
	for key, value := range m1 {
		fmt.Printf("%s=%d\n", key, value)
	}
}
