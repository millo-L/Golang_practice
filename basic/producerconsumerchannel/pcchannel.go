package main

import "fmt"

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100
}

func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

func sum(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a + b
	}()
	return out
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	d := sum(1, 2)

	fmt.Println(<-d)

	fmt.Scanln()
}
