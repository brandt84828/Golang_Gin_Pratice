
package main

import "fmt"

type Usb interface {
	//聲明兩個沒有實現的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

// 讓Phone實現Usb介面的方法
func (p Phone) Start() {
	fmt.Println("Phone start working")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop working")
}

func (p Phone) Call() {
	fmt.Println("Phone Call")
}

type Camera struct {
	name string
}

// 讓Camera實現Usb介面的方法
func (c Camera) Start() {
	fmt.Println("Camera start working")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop working")
}

// 電腦
type Computer struct {
}

// 寫一個方法working，接收一個Usb介面類型變數
// 只要是實現了Usb介面(所謂的實現Usb介面就是指實現了Usb介面聲明的所有方法)
func (c Computer) Working(usb Usb) {

	//如果usb是指向Phone結構體變數，則還需要調用call方法
	//使用類型斷言
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Start()
	usb.Stop()
}
func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"Google"}
	usbArr[1] = Phone{"Apple"}
	usbArr[2] = Camera{"Nikon"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
