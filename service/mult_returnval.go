package service

import (
	"fmt"
	"strings"
)

func StringConvmain() {

	fmt.Println(strings.IndexFunc("a", isA))

	var a = 3
	var b = 1
	var A, B, C int
	var A1, B1, C1 int
	A, B, C = Return1(a, b)
	fmt.Printf("非命名返回:{0},{1},{2}", A, B, C)
	Return2(a, b)
	fmt.Println(fmt.Printf("非命名返回:{0},{1},{2}", A1, B1, C1))

}

func isA(rune2 rune) bool {

	return rune2 == 'a'

}

func Return1(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func Return2(a, b int) (A1, B1, C1 int) {
	A1 = a + b
	B1 = a * b
	C1 = a - b
	return
}
