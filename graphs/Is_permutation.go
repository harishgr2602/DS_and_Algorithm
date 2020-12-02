package main

import "fmt"

func isPermutation(s1 string, s2 string) bool {
	count := make(map[byte]int)
	length := len(s1)
	if len(s2) != length {
		fmt.Println(s1, "&", s2, "are not permutaution")
		return false
	}
	for i := 0; i < length; i++ {
		ch := s1[i]
		count[ch]++
		ch = s2[i]
		count[ch]--
	}

	for i := 0; i < length; i++ {
		ch := s1[i]
		if count[ch] != 0 {
			fmt.Println(s1, "&", s2, "are not permutaution")
			return false
		}
	}
	fmt.Println(s1, "&", s2, "are permutaution")
	return true
}
