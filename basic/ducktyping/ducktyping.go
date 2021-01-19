package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("quack~!")
}

func (d Duck) feathers() {
	fmt.Println("duck has white and gray feathers")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("A person imitate a duck")
}

func (p Person) feathers() {
	fmt.Println("A person has duck's feathers")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)

	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}
}
