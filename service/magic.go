package service

import "fmt"

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
