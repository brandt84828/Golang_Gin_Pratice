// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) getSum() float64 {
	return calculator.Num1 + calculator.Num2
}

func main() {
	var calculator Calculator
	calculator.Num1 = 10.5
	calculator.Num2 = 20.8
	fmt.Print(calculator.getSum())
}
