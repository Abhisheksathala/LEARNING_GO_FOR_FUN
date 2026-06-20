package main

import "fmt"

func main() {
	view1 := 1000
	view2 := 2000
	totviews := view1 + view2 // 3000

	avgView := totviews / 2 // 1500

	fmt.Println(avgView)

	rating := 4.7 // 4.7

	finalRating := int(rating) // 4
	fmt.Println(finalRating)

	name := "abhishek"

	name = name + " sathala"
	fmt.Println(name) // abhishek sathala

	var name1 int = 10
	fmt.Println(name1) // 10

}
