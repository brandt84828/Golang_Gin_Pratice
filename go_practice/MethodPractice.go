// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) Speak() {
	fmt.Printf("%v is a good man. \n", person.Name)
}

func (person Person) Cal() {
	fmt.Printf("1+...+1000 = %v \n", (1+1000)/2*1000)
}

func (person Person) Cal2(n int) {
	fmt.Printf("1+...+%v = %v \n", n, (1+n)/2*n)
}

func (person Person) GetSum(n int, m int) {
	fmt.Printf("%v + %v = %v", n, m, n+m)
}
func main() {
	var person Person
	person.Name = "Tom"
	person.Speak()
	person.Cal()
	person.Cal2(20)
	person.GetSum(10, 20)
}
