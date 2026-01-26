package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 20,
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valC, okC := points["b"]
	fmt.Println(valC, okC)
	valD, okD := points["c"]
	fmt.Println(valD, okD)

	prices := map[string]int{
		"xyx": 500,
		"dzy": 500,
	}

	total := 0

	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println(total)

}
