package implementations

import "fmt"

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack!")
}

type Mute struct{}

func (q *Mute) Quack() {
	fmt.Println("<<silence>>")
}

type Squeak struct{}

func (q *Squeak) Quack() {
	fmt.Println("Squeak!")
}
