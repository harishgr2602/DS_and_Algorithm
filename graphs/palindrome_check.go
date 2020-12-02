package main

import "fmt"

func isPalindrome(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j && str[i] == str[j] {
		i++
		j--
	}
	if i < j {
		fmt.Println("String is not a palindrome")
		return false
	}
	fmt.Println("String is a palindrome")
	return true
}
