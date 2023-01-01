// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxIndex := 0
	for index, val := range intArr {
		if val > maxVal {
			maxVal = val
			maxIndex = index
		}
	}

	fmt.Printf("%v %v", maxIndex, maxVal)

}
