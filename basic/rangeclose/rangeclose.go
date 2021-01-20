package main

import "fmt"

func main() {
	c := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	a, ok := <-c
	fmt.Println(a, ok)
}
