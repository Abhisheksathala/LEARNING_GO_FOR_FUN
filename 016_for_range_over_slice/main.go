package main

import "fmt"

func main() {
	views := []int{10, 20, 45, 50, 60}
	// for range
	total := 0
	for i, v := range views {
		fmt.Println("day", i, "views", v) // it print the index and the value
		// but we can also print only the value by using _
		// for i, v := range views {
		// 	fmt.Println("views", i, v)
		// }
		// but we can also print only the index by using _
		// for i, _ := range views {
		// 	fmt.Println("day", i)
		// }

		total = total + v
	}
	fmt.Println(total)
}
