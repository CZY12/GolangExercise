package service

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (receiver Car) Start() {
	fmt.Println("The car's Start")
}

func (receiver Car) Stop() {
	fmt.Println("The car's Stop")
}
func (receiver *Car) SetWheel(num int) {
	receiver.wheelCount = num
}

func (receiver *Car) NumberOfWheels() {
	fmt.Println("The car's wheelCount is: %n", receiver.wheelCount)
}

func (receiver Mercedes) SayHiToMerkel() {
	fmt.Println("Hi,Merkel")
}
