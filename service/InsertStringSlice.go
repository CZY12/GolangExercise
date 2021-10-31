package service

/**
将切片插入到另一个切片的指定位置。
*/
func InsertStringSlice(newSlice, str []string, index int) (ns []string) {

	ns = append(ns, str[:index]...)
	ns = append(ns, newSlice[:]...)
	ns = append(ns, str[index:]...)
	return

}

/**
接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
*/
func CutStr(str string, index int) (start, end string) {
	startS := str[:index]
	endS := str[index:]

	return startS, endS
}

/**
反转字符串
*/
func ReverseString(str string) (s string) {
	bt := []byte(str)
	for i := 0; i < len(bt)/2; i++ {
		a := bt[i]
		if i < len(bt)-1-i {
			bt[i] = bt[len(bt)-1-i]
			bt[len(bt)-1-i] = a
		} else {
			break
		}
	}
	s = string(bt)
	return
}

/**
遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
*/

func SearchNoEq(str []string) (noEq []string) {
	noEq = make([]string, 0)
	for i, s := range str {
		if i+1 < len(str) {
			if s != str[i+1] {
				noEq = append(noEq, str[i+1])
			}
		} else {
			break
		}
	}
	return
}

/**
冒泡排序
*/
func BubbleSort(s []int) (e []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] < s[j] {
				var a = s[i]
				s[i] = s[j]
				s[j] = a
			} else {
				continue
			}
		}
	}
	return s
}

func MapFunc(f func(int) int, is []int) (is1 []int) {
	is1 = is
	for i, num := range is {
		is1[i] = f(num)
	}
	return
}

func MapFunc1(i int) (n int) {
	n = i * 10
	return
}
