package main

import "fmt"

func main() {
	// make([],length,capacity)
	score := make([]int, 0, 5)

	fmt.Println(score, len(score), cap(score))
	score = append(score, 10, 12, 14, 15, 16)
	fmt.Println(score, len(score), cap(score))

	todos := []string{"do yt", "workout everyday"}
	more := []string{"learn golanf..."}

	todos = append(todos, more...)

	fmt.Println(todos)
}
