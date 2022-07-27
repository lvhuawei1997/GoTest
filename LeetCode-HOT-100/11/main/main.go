package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxCap := 0
	for left < right {
		cur := (right - left) * min(height[left], height[right])
		if cur > maxCap {
			maxCap = cur
		}
		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}
	return maxCap
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	num := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(num))
}
