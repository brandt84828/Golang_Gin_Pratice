// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(fbn(10))
}

func fbn(n int) []uint64 {

	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
