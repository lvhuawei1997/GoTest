package main

import (
	"container/list"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	n := len(heights) - 1
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := 0; i <= n; i++ {
		right[i] = n + 1
	}
	stack := list.New()
	res := 0
	for i := 0; i <= n; i++ {
		for stack.Len() > 0 && heights[i] <= heights[stack.Back().Value.(int)] {
			right[stack.Back().Value.(int)] = i
			stack.Remove(stack.Back())
		}
		if stack.Len() == 0 {
			left[i] = -1
		} else {
			left[i] = stack.Back().Value.(int)
		}
		stack.PushBack(i)
	}
	//stack = list.New()
	//for i := n; i >= 0; i-- {
	//	for stack.Len() > 0 && heights[i] <= heights[stack.Back().Value.(int)] {
	//		stack.Remove(stack.Back())
	//	}
	//	if stack.Len() == 0 {
	//		right[i] = n + 1
	//	} else {
	//		right[i] = stack.Back().Value.(int)
	//	}
	//	stack.PushBack(i)
	//}
	for i := 0; i <= n; i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}
	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3}))
}
