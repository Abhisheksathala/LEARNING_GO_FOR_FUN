package main

import "fmt"

func main() {

	// slices in the GO most common collection Type
	// these are like dynamic arrays
	// most imporatnat in the GO

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

	s = s[1:] // now s will be [2 3 4] but why it cut the 1??
	// it remove the first elment from the slice

	s = s[:len(s)-1] // now s will be [1 2 3] but why it cut the 4??
	// it remove the last elment from the slice

	fmt.Println(s)
}
