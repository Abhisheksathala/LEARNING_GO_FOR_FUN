package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sumproduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func divied(a int, b int) (john int, sangam int) {
	john = a / b
	sangam = a + b
	return
}

func main() {
	sum1 :=
		add(1, 5)
	fmt.Println(sum1)
	sum2, sum3 := sumproduct(3, 5)
	// sum4, _ := sumproduct(3, 5)  u can one value by useing single value bro
	s1, s2 := divied(10, 10)
	fmt.Println(sum2, sum3)
	fmt.Println(s1, s2)
}
