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

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
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

	//Monster實現了AInterface和BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
