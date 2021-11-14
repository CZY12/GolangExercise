package service

type Employee struct {
	salary float32
}

/**
根据百分比增加薪水
*/
func (receiver Employee) GiveRaise(num float32) float32 {
	return (1 + num) * receiver.salary
}

func (receiver *Employee) SetSalary(num float32) {
	receiver.salary = num
}
