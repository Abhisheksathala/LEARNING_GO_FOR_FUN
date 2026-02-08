package main

import "fmt"

func main() {
	// common collection type
	// dynamic and can grow
	// []type{....}
	result := []string{"sangam", "abhishek"}

	fmt.Println(result, result[0], result[len(result)-1])

	result[1] = "abhi"

	var nums []int

	nums = append(nums, 10, 20, 80) // u can do it mulitipul times thats it
	fmt.Println(nums)

	s := []int{1, 2, 3, 4}

	s = s[1:]
	fmt.Println(s)
}
