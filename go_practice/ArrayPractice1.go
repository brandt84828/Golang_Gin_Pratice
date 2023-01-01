// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var byteArr [26]byte

	for i := 0; i < 26; i++ {
		byteArr[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c", byteArr[i])
	}

}
