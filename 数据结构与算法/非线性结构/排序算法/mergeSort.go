// 并归排序
package main

import "fmt"

// MergeSort 实现
func MergeSort(array []int) []int {
	Num := len(array)
	if Num < 2 {
		return array
	}

	key := Num / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}
		} else {
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}

func main() {
	array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println(array)
	array = MergeSort(array)
	fmt.Println(array)
}

/*
// 并归排序
package main

import (
	"fmt"
)

func merge(data []int) []int {
	sum := len(data)
	if sum <= 1 {
		return data
	}

	left := data[0 : sum/2]
	//fmt.Println(left)
	lSize := len(left)
	if lSize >= 2 {
		left = merge(left)
	}

	right := data[sum/2:]
	//fmt.Println(right)
	rSize := len(right)
	if rSize >= 2 {
		right = merge(right)
	}

	j, t := 0, 0
	arr := make([]int, sum)
	fmt.Println(left, right, data)
	for i := 0; i < sum; i++ {
		if j < lSize && t < rSize {
			if left[j] <= right[t] {
				arr[i] = left[j]
				j++
			} else {
				arr[i] = right[t]
				t++
			}
		} else if j >= lSize {
			arr[i] = right[t]
			t++
		} else if t >= rSize {
			arr[i] = left[j]
			j++
		}
	}
	return arr
}

func main() {

	testArr := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println(testArr)

	ResArr := merge(testArr)
	fmt.Println(ResArr)
}
*/
