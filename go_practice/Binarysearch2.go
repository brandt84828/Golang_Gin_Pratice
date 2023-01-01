// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	Search(arr, 0, 6, 88)

}

func Search(arr [6]int, leftindex int, rightindex int, findVal int) {

	if leftindex > rightindex {
		fmt.Println("Can't not find")
		return
	}

	mid := (leftindex + rightindex) / 2

	if arr[mid] > findVal {
		Search(arr, leftindex, mid-1, findVal)
	} else if arr[mid] < findVal {
		Search(arr, mid+1, rightindex, findVal)
	} else {
		fmt.Printf("Got it! index: %v", mid)
	}
}
