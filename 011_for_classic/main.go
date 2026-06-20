package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("jai sri ram", i)
	}

	// the output will be
	// first it checks i := 0; i < 5 and 0 < 5 is true then it print
	// then i++ i becomes 1 and 1 < 5 is true then it print
	// then i++ i becomes 2 and 2 < 5 is true then it print
	// then i++ i becomes 3 and 3 < 5 is true then it print
	// then i++ i becomes 4 and 4 < 5 is true then it print
	// then i++ i becomes 5 and 5 < 5 is false then it stop

	n := 10
	sum := 0

	for i := 0; i <= n; i++ {
		sum = sum + i
	}

	// the output will be
	// i := 0; i <= n; i++ and 0 <= 10 is true then it add
	// i++ i becomes 1 and 1 <= 10 is true then it add
	// i++ i becomes 2 and 2 <= 10 is true then it add
	// i++ i becomes 3 and 3 <= 10 is true then it add
	// i++ i becomes 4 and 4 <= 10 is true then it add
	// i++ i becomes 5 and 5 <= 10 is true then it add
	// i++ i becomes 6 and 6 <= 10 is true then it add
	// i++ i becomes 7 and 7 <= 10 is true then it add
	// i++ i becomes 8 and 8 <= 10 is true then it add
	// i++ i becomes 9 and 9 <= 10 is true then it add
	// i++ i becomes 10 and 10 <= 10 is true then it add
	// i++ i becomes 11 and 11 <= 10 is false then it stop

	fmt.Println(sum)

}
