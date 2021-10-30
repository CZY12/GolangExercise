package service

/**
将从 start 到 end 索引的元素从切片 中移除。
*/
func RemoveStringSlice(str []string, start, end int) (ns []string) {
	for i, num := range str {
		if i >= start && i < end {
			continue
		}
		ns = append(ns, num)
	}
	return
}
