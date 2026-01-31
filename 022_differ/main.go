package main

import (
	"errors"
	"fmt"
)

func main() {
	// defer resp.bosy.close()

}

func doWork(success bool) error {
	// resource related
	// started message -> resource aquried
	// cleanup message -> resource released
	fmt.Println("start: resource aquired")
	defer fmt.Println("cleanup: resource relesed")
	if !success {
		return errors.New("somthing waent wrong . i am returning early ")
	}
	fmt.Println("work: doing somthing imp")
	fmt.Println("work: this work is done")
	return nil
}

// 2:5:18 here i stoped the video of golang
