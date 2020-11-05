package main

import (
	"fmt"
	"github.com/priyeshbhaskar/learngo/channels"
	"github.com/priyeshbhaskar/learngo/hello"
)


func main() {

	fmt.Println(hello.SayHello())
	c := channels.GenNumber([]int{2, 3, 4, 5})

	// Consume the output.
	// Print 2,3,4,5
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
