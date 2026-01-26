package main

import (
	"fmt"
	"strconv"
)

func run() error {
	input := "3"
	level, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("selected level", level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value,err)
	// nil error -> success
	// non nil -> fail

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("there is an error")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("levelk must be 1 or 5 between")
	}

	return n, nil

}

func main() {
	//  go dont use exceptions for normal failures
	// functions -> return errors as normal return values

	// val , error := somting()
	run()
}
