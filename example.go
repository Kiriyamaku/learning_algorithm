package main
import "fmt"

//stack
func main(){

	stack := make([]int,0)
	stack = append(stack, 10)
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("test")
	fmt.Println(v)
}
