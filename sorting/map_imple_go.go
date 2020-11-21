package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Println("[", key, "->", val, "]")
	}
	fmt.Println("Apple Price:", m["Apple"])
	delete(m, "Apple")
	fmt.Println("Apple Price:", m["Apple"])

	v, ok := m["Apple"]
	fmt.Println("Apple Price:", v, "Present", ok)

	v2, ok2 := m["Banana"]
	fmt.Println("Banana Price:", v2, "Present", ok2)
}
