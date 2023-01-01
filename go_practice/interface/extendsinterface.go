package main

import _ "fmt"

type AInterface interface {
	test01()
}

type BInterface interface {
	test02()
}

type CInterface interface {
	AInterface
	BInterface
	test03()
}

// 如果需要實現C這個interface，就需要將AInterface和BInterface的方法都實現
type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

func main() {

	var stu Stu
	var a CInterface = stu
	a.test01()

}
