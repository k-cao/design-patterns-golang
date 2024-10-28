package ducks

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/strategy/implementations"
)

type WoodenDuck struct {
	Duck
}

func NewWoodenDuck() *WoodenDuck {
	return &WoodenDuck{
		Duck: Duck{
			flyBehavior:   &implementations.FlyNoWay{},
			quackBehavior: &implementations.Mute{},
			displayImplementation: func() {
				fmt.Println("I'm a wooden duck")
			},
		},
	}
}
