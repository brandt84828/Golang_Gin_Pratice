// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var intArr [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	avg := 0.0

	for _, val := range intArr {
		sum += val
	}
	avg = float64(sum) / float64(len(intArr))

	fmt.Printf("%v %v", sum, avg)

}
