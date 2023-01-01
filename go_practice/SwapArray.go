// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var intArr [5]int

	//for every time get different random number
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}

	fmt.Println(intArr)

	tmp := 0
	for i := 0; i < len(intArr)/2; i++ {
		tmp = intArr[i]
		intArr[i] = intArr[len(intArr)-i-1]
		intArr[len(intArr)-i-1] = tmp
	}

	fmt.Println(intArr)

}
