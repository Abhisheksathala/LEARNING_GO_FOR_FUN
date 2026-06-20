package main

import (
	"fmt"
)

func main() {
	games := 10
	pricegames := 599
	if total := games * pricegames; total >= 6000 {
		fmt.Println("good man its")
	} else {
		fmt.Println("somthing")
	}
}
