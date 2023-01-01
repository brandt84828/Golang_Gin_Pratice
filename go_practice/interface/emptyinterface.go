package main

import (
	"fmt"
)

type Stu struct {
}

type T interface {
}


//interface默認是一個指針(引用類型)，如果沒有初始化就使用會輸出nil
//空的interface沒有任何方法，所以所有類型都實現了空的interface，可以把任何一個變數賦予給空的interface
func main() {

	var stu Stu
	var t T = stu
	fmt.Println(t)

	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2)

}
