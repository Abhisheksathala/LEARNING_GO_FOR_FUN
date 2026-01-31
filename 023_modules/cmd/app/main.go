package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main() {
	mes1 := greet.Hello("abhishek")
	fmt.Println(mes1)
}
