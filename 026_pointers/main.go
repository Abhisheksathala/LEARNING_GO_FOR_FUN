package main

import "fmt"

func main() {
	// store the memory address of any val
	// &x -> address of x (makes a pointer)
	// *p -> defer

	score := 10

	fmt.Println("bef", score)
	addscrol(&score)
	fmt.Println("aft", score)
}

func addscrol(scroe *int) {
	*scroe = *scroe + 5
}
