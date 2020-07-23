package MyTest

import (
	"fmt"
)

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

type animal interface {
	say() string
	eat() string
}

func (d dog) say() (word string) {
	word = "wang"
	return word
}

func (d dog) eat() (food string) {
	food = "bone"
	return food
}

func (c cat) say() (word string) {
	word = "wow"
	return word
}

func (c cat) eat() (food string) {
	food = "fish"
	return food
}

func tt() {
	tom := new(cat)
	fmt.Println(animal.eat(tom))
}
