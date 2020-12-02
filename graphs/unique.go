package main

import "fmt"

func isUniqueChar(str string) bool {
	mp := make(map[byte]int)
	size := len(str)
	for i := 0; i < size; i++ {
		c := str[i]
		if mp[c] != 0 {
			fmt.Println("Duplicate Deleted!")
			return false
		}
		mp[c] = 1
	}
	fmt.Println("No Duplicate Detedted!")
	return true
}
