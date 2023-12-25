package main

import (
	"fmt"
	"slices"
)

type book string

func (b book) printTitle() {
	fmt.Println(b)
}

func twoSum(nums []int, target int) []int {
	var result []int

	for i, value := range nums {
		diff := target - value
		fmt.Println(diff)
		// if nums[i] == diff {
		// 	continue
		// }
		var s []int = nums[i+1:]
		fmt.Println(s)
		if slices.Contains(s, diff) {
			idx := slices.Index(nums, diff)
			result = append(result, i, idx)
			break
		}
	}
	return result
}

func main() {
	// var bookTitle book = "Game Of Thrones"
	// bookTitle.printTitle()
	nums := []int{3, 3} //[3,2,4] [2, 7, 11, 15]
	target := 6

	result := twoSum(nums, target)
	fmt.Println(result)

}

/*
	receiver function with retune type
*/
// type laptopSize float64

// func (this laptopSize) getSizeOfLaptop() laptopSize {
// 	return this
// }
