package main

import (
	"fmt"
)

func main() {
	s := "Hello world!!!"
	r := []rune(s)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
}
