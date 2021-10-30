package main

import (
	"CzyCoding/service"
	"fmt"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	a := rand.Int()
	//	fmt.Printf("%d / ",a)
	//}
	//fmt.Println()
	//for i := 0; i < 5; i++ {
	//	r := rand.Intn(8)
	//	fmt.Printf("%d / ",r)
	//}
	//fmt.Println()
	//times := int64(time.Now().Nanosecond())
	//fmt.Printf("%v",times)
	//fmt.Println()
	//rand.Seed(times)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%2.2f /",100*rand.Float32())
	//
	//}
	var str = []string{"1", "2", "3", "7"}
	var newSlice = []string{"4", "5", "6"}

	var ns = service.InsertStringSlice(newSlice, str, 3)

	for i, num := range ns {
		fmt.Println("%d is %n", i, num)
	}
	var i1 = 5
	var intP *int
	intP = &i1
	fmt.Printf("An integer: %d, it's lpcation in memory: %p\n", i1, &i1)

	fmt.Printf("int : %p", intP)

	//var a = 3
	//var b = 1
	//var A, B, C int
	//var A1, B1, C1 int
	////A, B, C = Return1(a, b)
	//fmt.Printf("非命名返回:{0},{1},{2}", A, B, C)
	//A1, B1, C1 = Return2(a, b)
	//fmt.Println(fmt.Printf("非命名返回:{0},{1},{2}", A1, B1, C1))

}
