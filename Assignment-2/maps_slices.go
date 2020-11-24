package main
 
import "fmt"

func main() { 
    m_a_p := map[int]string{ 
  
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    for id, pet := range m_a_p { 
  
        fmt.Println(id, pet) 
    } 
} 