// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print(len float64, width float64) float64 {
	return len * width
}

func main() {
	var mu MethodUtils
	area := mu.Print(10, 8)
	fmt.Print("Area=", area)
}
