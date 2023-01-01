// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var arr = [5]int{14, 655, 631, 2, 34}
	fmt.Println(arr)

	Bubblesort(&arr)

	fmt.Println(arr)

}

func Bubblesort(arr *[5]int) {

	temp := 0
	for i := 1; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {

			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}

		}

	}

}
