package main

import (
	"fmt"
)

func IsBalancedParanthesis(expn string) bool {
	stk := new(Stack)
	for _, ch := range expn {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}

func main() {
	expn := "{()}[]"
	value := IsBalancedParanthesis(expn)
	fmt.Println("Given Expn:", expn)
	fmt.Println("Result after Paranthesis is matched:", value)
}
