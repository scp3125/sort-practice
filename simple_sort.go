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

// 归并排序
func MergeSort(list []int) {
	mid := len(list) / 2

	MergeSort(list[:mid])
	MergeSort(list[mid+1:])
	merge(list[:mid], list[mid+1:])
}

func merge(list1 []int, list2 []int) {
	idx1 := 0
	idx2 := 0

	n1 := len(list1)
	n2 := len(list2)
	for idx1 < n1 || idx2 < n2 {

	}
}
