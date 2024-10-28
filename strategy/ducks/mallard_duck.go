package ducks

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/strategy/implementations"
)

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		Duck: Duck{
			flyBehavior:   &implementations.FlyWithWings{},
			quackBehavior: &implementations.Quack{},
			displayImplementation: func() {
				fmt.Println("I'm a mallard duck")
			},
		},
	}
}
