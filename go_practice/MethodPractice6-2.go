// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calculator.Num1 + calculator.Num2
	case '-':
		res = calculator.Num1 - calculator.Num2
	case '*':
		res = calculator.Num1 * calculator.Num2
	case '/':
		res = calculator.Num1 / calculator.Num2
	default:
		fmt.Print("Wrong Operator")
	}

	return res
}

func main() {
	var calculator Calculator
	calculator.Num1 = 10.5
	calculator.Num2 = 20.8
	fmt.Print(calculator.getRes('*'))
}
