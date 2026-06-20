package main

import "fmt"

func main() {
	// make([],length,capacity)
	score := make([]int, 0, 5)
	// make ? what is it why we use ?

	fmt.Println(score, len(score), cap(score))
	score = append(score, 10, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)
	fmt.Println(score, len(score), cap(score))

	todos := []string{"do yt", "workout everyday"}
	more := []string{"learn golanf..."}

	// this line does what ?
	todos = append(todos, more...) // this line appends the more slice to the todos slice not only this it also resizes the slice if needed

	fmt.Println(todos)
}
