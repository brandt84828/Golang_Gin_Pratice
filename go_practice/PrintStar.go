// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var totalLevel int = 9

	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

			//fmt.Print("*")
		}
		fmt.Println()
	}
}
