package main

import "fmt"

func main() {
	var city string
	city = "abhishek"

	fmt.Println("city : ", city)

	// := inefrance is that langae decides its type (runtime)

	sub := 5000
	sub = sub + 1000
	fmt.Println("sub :", sub) // 6000

	like, comment := 100, "1000"
	fmt.Println("like :", like, "comment :", comment) // 100 "1000"

}
