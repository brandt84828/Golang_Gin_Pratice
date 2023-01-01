// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print(m int, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println("")
	}
}

func main() {
	var mu MethodUtils
	mu.Print(10, 8, "@")
}
