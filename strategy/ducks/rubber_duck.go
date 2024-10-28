package ducks

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/strategy/implementations"
)

type RubberDuck struct {
	Duck
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{
		Duck: Duck{
			flyBehavior:   &implementations.FlyNoWay{},
			quackBehavior: &implementations.Squeak{},
			displayImplementation: func() {
				fmt.Println("I'm a rubber duck")
			},
		},
	}
}
