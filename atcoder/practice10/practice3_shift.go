package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var i int
	fmt.Scanf("%d", &n)

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()

	fmt.Printf("%d\n", strings.Count(s, "1"))
}
