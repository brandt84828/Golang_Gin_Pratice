
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type MethodUtils struct {
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Print("even number")
	} else {
		fmt.Print("odd number")
	}
}

func main() {
	var mu MethodUtils
	mu.JudgeNum(20)
}
