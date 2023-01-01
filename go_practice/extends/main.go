package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

type Pupil struct {
	Student //嵌入匿名結構體
}

type Graduate struct {
	Student //嵌入匿名結構體
}

// show
func (stu *Student) ShowInfo() {
	fmt.Printf("Name=%v Age=%v Score=%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// method(share)
func (stu *Student) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func (p *Pupil) testing() {
	fmt.Println("Pupil testing...")
}

func (p *Graduate) testing() {
	fmt.Println("Graduate testing...")
}

func main() {

	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.getSum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.getSum(10, 20))
}
