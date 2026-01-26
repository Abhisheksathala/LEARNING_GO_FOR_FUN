package main

import (
	"fmt"
)

func main() {
	games := 3
	pricegames := 59

	if total := games * pricegames; total >= 100 {
		fmt.Println("good man its")
	} else {
		fmt.Println("somthing")
	}

}
