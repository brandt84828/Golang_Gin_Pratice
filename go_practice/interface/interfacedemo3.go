package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

func main() {

	//error 不能直接實例化interface
	//var a AInterface

	var stu Stu //結構體變數，實現了Say()-->實現了AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()
}
