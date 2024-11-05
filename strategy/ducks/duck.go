package ducks

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/strategy/behaviors"
)

type Duck struct {
	flyBehavior   behaviors.FlyBehavior
	quackBehavior behaviors.QuackBehavior

	displayImplementation func()
}

func (d *Duck) Fly() {
	d.flyBehavior.Fly()
}

func (d *Duck) Quack() {
	d.quackBehavior.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("All ducks can swim")
}

func (d *Duck) Display() {
	d.displayImplementation()
}

func (d *Duck) SetFly(b behaviors.FlyBehavior) {
	d.flyBehavior = b
}

func (d *Duck) SetQuack(b behaviors.QuackBehavior) {
	d.quackBehavior = b
}
