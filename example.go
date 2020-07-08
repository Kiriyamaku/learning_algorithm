package main
import "fmt"

//stack
func main(){

	//stack
	stack := make([]int,0)
	//push
	stack = append(stack, 10)
	//pop
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	//check empty
	len(stack) == 0
	// fmt.Printf("test")
	// fmt.Println(v)

	//queue
	queue := make([]int,0)
	//enqueue
	queue = append(queue,10)
	//dequeue
	queue = queue[1:]
	// get dequeue item
	v := queue[0]
	// check empty
	len(queue)==0

	//dict
	m:=make(map[string]int)
	//add item
	m["test"]=1
	//delete item
	delete(m,"test")
	//for range
	for k,v := range m{
		println(k,v)
	}

	sort.Ints([]int{})

	// use copy to delete item in 
	copy(a[i:],a[i+1:])
	a=a[:len(a)-1]

	//convert type
	s :="12345"
	m:=s[0]-'1'
	fmt.Printf("%v",m)
}
