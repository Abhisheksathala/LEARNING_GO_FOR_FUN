package main

import "fmt"

// struct group related fileds into one type
type user struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u := user{Name: "abhi", ID: 1, Email: "abhishek@gamil.com", Age: 19}
	fmt.Println("befo", u)
	fmt.Println("********")
	u.Birthday()
	fmt.Println("after", u)
}

func (u *user) Birthday() {
	u.Age++
}
