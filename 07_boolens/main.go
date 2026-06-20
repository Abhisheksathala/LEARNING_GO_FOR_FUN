package main

import (
	"fmt"
)

func main() {
	isLogged := true
	fmt.Println(isLogged) // false

	isadmin := false
	hasSubscribtion := true

	canopenDashboard := isLogged && hasSubscribtion
	candeletPost := isadmin || isLogged && hasSubscribtion

	fmt.Println("canopenDashboard", canopenDashboard)
	fmt.Println("candeletPost", candeletPost)

	if isLogged {
		fmt.Println("User is logged in")
	} else {
		fmt.Println("User is not logged in")
	}

	if !isLogged {
		fmt.Println("User is not logged in")
	} else {
		fmt.Println("User is logged in")
	}
}
