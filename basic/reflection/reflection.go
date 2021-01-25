package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var s string = "Hello, World!"
	fmt.Println(reflect.TypeOf(s))

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f))

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data))

	var f64 float64 = 1.3
	t := reflect.TypeOf(f64)
	v := reflect.ValueOf(f64)

	fmt.Println(t.Name())                    // float64: 자료형 이름 출력
	fmt.Println(t.Size())                    // 8: 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // true: 자료형 종류를 알아내어
	// reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64) // false: 자료형 종류를 알아내어 reflect.Int64와 비교

	fmt.Println(v.Type())                    // float64: 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // true: 값이 담긴 변수의 자료형 종류를
	// 알아내어 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64) // false: 값이 담긴 변수의 자료형 종류를
	// 알아내어 reflect.Int64와 비교
	fmt.Println(v.Float()) // 1.3: 값을 실수형으로 출력
}
