package service

import (
	"fmt"
)

func main() {

	// 判断字符串s前缀是否为 str
	//	var s = "abcd"
	//	var str = "a"
	//
	//	fmt.Printf("%s is %s 's prefix: %t", str, s, isPre(s, str))
	//
	//	var i = 1
	//	switch {
	//	case i <= 3:
	//		fmt.Printf("春季")
	//		fallthrough
	//	case i <= 6:
	//		fmt.Printf("夏季")
	//		fallthrough
	//	case i <= 9:
	//		fmt.Printf("秋季")
	//		fallthrough
	//	case i <= 12:
	//		fmt.Printf("冬季")
	//		fallthrough
	//	default:
	//		fmt.Printf("输入异常")
	//	}
	//
	//STAR:
	//	fmt.Println(i)
	//	i++
	//	if i < 15 {
	//		goto STAR
	//	}
	//
	//	var str1 = "GGGGGGGGGGGGGGGGGGGGGGGGG"
	//	for i := 1; i < len(str1); i++ {
	//		connent := str1[0 : i-1]
	//		fmt.Println(connent)
	//	}

	//for i:= 1; i<=100;i++ {
	//	switch  {
	//	case i%3 == 0 && i% 5 == 0:
	//		fmt.Println("FizzBuzz")
	//	case i%3 == 0:
	//		fmt.Println("Fizz")
	//	case i% 5 == 0:
	//		fmt.Println("Buzz")
	//	default:
	//		fmt.Println(i)
	//	}
	//}

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func isPre(s, pre string) bool {
	if len(s) > len(pre) {
		return s[0:len(pre)] == pre
	}
	return false
}
