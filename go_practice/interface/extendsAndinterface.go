package main

import "fmt"

type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "Climbing...")
}

type LittleMonkey struct {
	Monkey //extends
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

// 讓LittleMonkey實現interface
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "Flying")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "Swimming")
}

func main() {

	monkey := LittleMonkey{
		Monkey{
			Name: "Wu",
		},
	}
	var a BirdAble = &monkey
	monkey.climbing()
	a.Flying()

	var b FishAble = &monkey
	b.Swimming()
}
