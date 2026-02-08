package main

import "fmt"

func main() {
	var marks [3]int
	// in this go  arrays are in the fix size and never grow like in the js
	marks[0] = 10
	marks[1] = 20

	prime := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(prime)
}
