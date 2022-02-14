package main

// 冒泡排序
func BubbleSort(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for idx := 0; idx < n-1-i; idx++ {
			if list[idx] > list[idx+1] {
				list[idx], list[idx+1] = list[idx+1], list[idx]
			}
		}
	}
}

// 选择排序
func SelectionSort(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		minIdx := i
		for idx := i + 1; idx < n; idx++ {
			if list[idx] < list[minIdx] {
				minIdx = idx
			}
		}
		list[minIdx], list[i] = list[i], list[minIdx]
	}
}

// 插入排序
func InsertionSort(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for idx := i; idx > 0; idx-- {
			if idx-1 < 0 {
				break
			}
			if list[idx] < list[idx-1] {
				list[idx], list[idx-1] = list[idx-1], list[idx]
			} else {
				break
			}
		}
	}
}
