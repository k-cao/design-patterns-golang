package main

/*
Pattern: Strategy
* Defines family - algos
* Encapsulates each one
* Make them interchangeable
* Lets algo vary independently from clients which use it
*/

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/strategy/ducks"
	"github.com/k-cao/design-patterns-golang/strategy/implementations"
)

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

	fmt.Println("Runtime behavioral changes below:")

	// Change behavior at runtime.
	dynamicDuck := ducks.Duck{}
	dynamicDuck.SetFly(&implementations.FlyWithWings{})
	dynamicDuck.SetQuack(&implementations.Squeak{})

	dynamicDuck.Quack()
	dynamicDuck.Fly()
	dynamicDuck.Swim()

	dynamicDuck.SetFly(&implementations.FlyNoWay{})
	dynamicDuck.SetQuack(&implementations.Mute{})

	dynamicDuck.Quack()
	dynamicDuck.Fly()
	dynamicDuck.Swim()
}
