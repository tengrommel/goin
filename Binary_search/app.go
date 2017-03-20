package main

import (
	"fmt"
	"time"
)

// 递归
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex{
		return -1
	}
	mid := int((lowIndex + highIndex)/2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	}else if array[mid] < target{
		return binarySearch(array, target, mid+1, highIndex)
	}else {
		return mid
	}
}
// 迭代
func iterBinarySearch(array []int, target int, lowIndex int,  highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex  {
		mid = int((startIndex + endIndex)/2)
		if array[mid] > target{
			endIndex = mid
		} else if array[mid] < target{
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	slice := []int{1,2,3,4,5,6,6,10,15,15,17,18,19,30}
	now := time.Now()
	find := binarySearch(slice, 10, 0, len(slice)-1)
	end := time.Now()
	run := end.Sub(now)
	fmt.Printf("Found the element by recursive binary search at index %d %v\n", find, run)
	now = time.Now()
	iterfind := iterBinarySearch(slice, 10, 0, len(slice)-1)
	end = time.Now()
	run = end.Sub(now)
	fmt.Printf("Found the element by recursive binary search at index %d  %v\n", iterfind, run)
}
