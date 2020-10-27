package main

import (
	"fmt"
	"github.com/priyeshbhaskar/learngo/hello"
	"github.com/priyeshbhaskar/learngo/channels"
)


func main() {

	fmt.Println(hello.SayHello())

	c := make(chan int)
	c = Gen([]int{2, 3, 4, 5})

	// Consume the output.
	// Print 2,3,4,5
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
