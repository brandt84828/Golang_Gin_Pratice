package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A Say OK", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A    //extends
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B Say OK", b.Name)
}

func main() {
	var b B
	b.A.Name = "Tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()

	//can simplify
	// If b has match attribute or method, it will call it directly. If not, it will get anonymous struct. Not find > report error.
	// If B and A have same attribute or method. It will call nearest att/method.
	b.Name = "smith"
	b.age = 20
	b.A.Name = "scott"
	b.SayOk()
	b.hello()
}
