package greet

import "strings"

// reuseable components

// exported com start with capital letter
func Hello(name string) string {

	clean := normalizename(name)

	return "hello," + clean

}

func normalizename(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "gust"
	}

	return strings.ToUpper(n)

}
