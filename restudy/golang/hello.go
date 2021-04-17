package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	sum, i := 0, 0

	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum)

	c := 'a'

	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c은(는) 숫자입니다.\n", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c은(는) 소문자입니다.\n", c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c은(는) 대문자입니다.\n", c)
	}

	type Rect struct {
		width  int
		height int
	}

	r := Rect{1, 2}
	fmt.Println(r.width*2 + r.height*2)

	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row, rowValue := range table {
		for col, colValue := range rowValue {
			if colValue == x {
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				// LOOP로 지정된 for 문을 빠져나옴
				break LOOP
			}

		}
	}

	f()
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
