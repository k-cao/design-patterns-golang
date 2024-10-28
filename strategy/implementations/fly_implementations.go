package implementations

import "fmt"

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("I'm flying!")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
