package main

import "fmt"

func main() {
	// map[keyType] valueType

	ages := map[string]int{
		"sangam": 65,
		"jone":   75,
	}

	fmt.Println(ages["sangam"], len(ages))

	// make(map[k]v)

	var score map[string]int //nil map

	fmt.Println(score, score["a"])

	score = make(map[string]int)

	score["math"] = 90
	score["s"] = 70

	fmt.Println(score)

	user := map[string]string{
		"u1": "sangam",
		"u2": "sangams",
		"u3": "sangamss",
	}

	fmt.Println(user)

	delete(user, "u1")
	fmt.Println(user)

}
