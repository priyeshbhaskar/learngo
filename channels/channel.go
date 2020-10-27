package channels

import "fmt"

func GenNumber(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	fmt.Println("return statement is called ")
	return out
}
