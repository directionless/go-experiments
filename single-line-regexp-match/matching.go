package main

import (
	"fmt"
	"regexp"
)

func main() {
	m1, _ := regexp.MatchString("[aeiou]", "hello")
	if m1 {
		fmt.Println("has vowels")
	}

	if m2, _ := regexp.MatchString("[aeiou]", "goodbye"); m2 {
		fmt.Println("has vowels")
	}

}
