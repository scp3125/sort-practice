package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func main() {
	Logarithm(InsertionSort)
}

// 对数器
func Logarithm(tagert func([]int)) {
	rand.Seed(time.Now().UnixNano())
	n := 1000
	for i := 0; i < n; i++ {
		var list []int
		for j := 0; j < n; j++ {
			list = append(list, rand.Intn(n))
		}
		list1 := make([]int, n)
		list2 := make([]int, n)
		copy(list1, list)
		copy(list2, list)
		sort.Ints(list1)
		tagert(list2)

		if !reflect.DeepEqual(list1, list2) {
			fmt.Println(list2)
			panic("不相等")
		}

		if i == n-1 {
			fmt.Println(list2)
		}
	}
	fmt.Println("测试成功")
}
