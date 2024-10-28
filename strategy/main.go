package main

/*
Pattern: Strategy
* Defines family - algos
* Encapsulates each one
* Make them interchangeable
* Lets algo vary independently from clients which use it
*/

import "github.com/k-cao/design-patterns-golang/strategy/ducks"

func main() {
	mallard := ducks.NewMallardDuck()
	rubberDuck := ducks.NewRubberDuck()
	woodenDuck := ducks.NewWoodenDuck()

	for _, duck := range []*ducks.Duck{
		&mallard.Duck, &rubberDuck.Duck, &woodenDuck.Duck,
	} {
		duck.Quack()
		duck.Fly()
		duck.Display()
		duck.Swim()
	}
}
