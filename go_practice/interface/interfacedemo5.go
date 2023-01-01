package main

import "fmt"

type Usb interface {
	Say()
}

type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	var u Usb = stu //錯誤，因為這邊是用stu類型，但定義的Say方法用的*Stu則是指針類型，Stu類型沒有實現Usb interface
	//正確的要改 var u Usb = &stu
	u.Say()
	fmt.Println("Here", u)
}
