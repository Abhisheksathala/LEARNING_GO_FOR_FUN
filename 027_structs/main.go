package main

import "fmt"

// struct group related fileds into one type
type user struct {
	ID    int
	Name  string
	Email string
}

func main() {
	ui := user{
		ID:    1,
		Name:  "abhishek",
		Email: "abhishek@gmail.com",
	}

	// ui.Email = "abhi@gmail.com"

	fmt.Println(ui)
}
