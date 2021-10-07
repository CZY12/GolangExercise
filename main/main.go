package main

import (
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
	var i1 = 5
	var intP *int
	intP = &i1
	fmt.Printf("An integer: %d, it's lpcation in memory: %p\n", i1, &i1)

	fmt.Printf("int : %p", intP)

}
