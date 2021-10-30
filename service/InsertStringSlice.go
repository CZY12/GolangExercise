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
