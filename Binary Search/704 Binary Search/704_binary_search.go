// package main

// import "fmt"

// func search(nums []int, target int) int {
// 	low := 0
// 	high := len(nums) - 1

// 	for low <= high {
// 		mid := (low + high) / 2

// 		if nums[mid] < target {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	if low == len(nums) || nums[low] != target {
// 		return false
// 	}

// 	return -1

// }

// func main() {
// 	nums := []int{-1, 0, 3, 5, 9, 12}
// 	fmt.Println(search(nums, 9))
// }

package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}
