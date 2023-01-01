// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
	var mu MethodUtils
	mu.Print()
}
