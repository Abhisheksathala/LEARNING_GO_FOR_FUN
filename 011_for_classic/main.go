package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("jai sri ram", i)
	}

	n := 10
	sum := 0

	for i := 0; i <= n; i++ {
		sum = sum + i
	}

	fmt.Println(sum)

}
