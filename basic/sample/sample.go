package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) introduce() {
	fmt.Println("애완동물의 이름은 " + a.name + "입니다.")
}

type Cat struct {
	// Animal 구조체의 자료형만 선언하면 상속과 같은 역할
	Animal
}

type Dog struct {
	// Animal 구조체의 자료형만 선언하면 상속과 같은 역할
	Animal
}

func main() {
	a := Animal{"꿀꿀이"}
	kitty := Cat{}
	puppy := Dog{}

	kitty.Animal.name = "나비"
	puppy.Animal.name = "뭉치"

	a.introduce()     // 내 애완동물의 이름은 꿀꿀이입니다.
	kitty.introduce() // 내 애완동물의 이름은 나비입니다.
	puppy.introduce() // 내 애완동물의 이름은 뭉치입니다.

	kitty.Animal.introduce() // 내 애완동물의 이름은 나비입니다.
	puppy.Animal.introduce() // 내 애완동물의 이름은 뭉치입니다.
}
