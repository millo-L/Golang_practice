package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello, World!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
