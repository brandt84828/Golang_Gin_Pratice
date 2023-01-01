package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok

	//error -> b = a
	
	//correct，assert，b:=a.(Point)就是類型斷言，表示判斷a是否是指向Point類型的變數，如果是就轉成Point類型並賦值給b，否則報錯。
	//類型斷言，由於interface是一般類型，不知道具體類型，就需要使用類型斷言
	b := a.(Point)
	fmt.Println(b)
}
