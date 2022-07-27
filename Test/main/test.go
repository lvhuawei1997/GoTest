package main

import (
	"fmt"
)

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := false // 有没有做过交换数据，没有表示有序
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func selectSort(nums []int) {
	min := 0
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
	}
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i
		for ; j > 0 && nums[j-1] > temp; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = temp
	}
}

func merge(left, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		res = append(res, left...)
	}
	if len(right) != 0 {
		res = append(res, right...)
	}
	return res
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	leftArr := mergeSort(nums[:mid])
	rightArr := mergeSort(nums[mid:])
	return merge(leftArr, rightArr)
}

func shellSort(nums []int) {
	length := len(nums)
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			pre := i - gap
			for pre >= 0 && nums[pre] > temp {
				nums[pre+gap] = nums[pre]
				pre -= gap
			}
			nums[pre+gap] = temp
		}
		gap /= 2
	}
}

func getMid(i, j, k int) int {
	if i > j && i < k {
		return i
	} else if j > i && j < k {
		return j
	} else {
		return k
	}
}

//func partition(nums []int, low, high int) int {
//	tmp := nums[low]
//	for low < high {
//		for low < high && nums[high] >= tmp {
//			high--
//		}
//		nums[low] = nums[high]
//		for low < high && nums[low] <= tmp {
//			low++
//		}
//		nums[high] = nums[low]
//	}
//	nums[low] = tmp
//	return low
//}
//
//func quickSort(nums []int, low, high int) {
//	if low < high {
//		mid := partition(nums, low, high)
//		quickSort(nums, low, mid-1)
//		quickSort(nums, mid+1, high)
//	}
//}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	mid := (start + end) / 2
	key := getMid(nums[start], nums[mid], nums[end])
	for i <= j {
		for nums[i] < key {
			i++
		}
		for nums[j] > key {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if start < j {
		quickSort(nums, start, j)
	}
	if i < end {
		quickSort(nums, i, end)
	}
}

func sift(nums []int, start, length int) {
	root := start
	child := root*2 + 1
	for child < length {
		if length-child > 1 && nums[child] < nums[child+1] {
			child++
		}
		if nums[root] < nums[child] {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}

func heapSort(nums []int) {
	length := len(nums)
	start := length/2 + 1
	end := length - 1
	for start >= 0 {
		sift(nums, start, length)
		start--
	}
	for end > 0 {
		nums[end], nums[0] = nums[0], nums[end]
		sift(nums, 0, end)
		end--
	}
}

func main() {
	nums := []int{4, 5, 6, 3, 2, 1}
	fmt.Println(nums)
	s := nums[:]
	s = append(s, 1)
	heapSort(nums)
	fmt.Println(nums)
}
