package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0
	for _, currunt := range nums {
		total = total + currunt
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
