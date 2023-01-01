package main

import (
	"encapsulation/model"
	"fmt"
)

func main() {
	p := model.NewPersion("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age = ", p.GetAge(), " sal = ", p.GetSal())
}
