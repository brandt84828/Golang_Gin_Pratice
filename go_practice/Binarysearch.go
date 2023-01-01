// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	arr := []int{1, 8, 10, 89, 1000, 1234}
	fmt.Println(Search(arr, 10))

}

func Search(nums []int, target int) int {

	mid := len(nums) / 2

	if nums[mid] == target {
		return mid
	}
	if len(nums) >= 1 {
		if nums[mid] < target {
			return Search(nums[mid:], target) + mid
		} else {
			return Search(nums[:mid], target)
		}
	}

	return -1
}
