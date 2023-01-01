package main

import "fmt"

type Usb interface {
	//聲明兩個沒有實現的方法
	Start()
	Stop()
}

type Usb2 interface {
	//interface所有方法都要實現，不然會錯誤
	Start()
	Stop()
	Test()
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
func (c Computer) Working(usb Usb2) {

	//透過usb介面變數調用start & stop
	usb.Start()
	usb.Stop()
}
func main() {
	//定義Usb interface array，可以存放Phone和Camear的結構體變數(多型陣列/多態數組)
	var usbArr [3]Usb

	usbArr[0] = Phone{"Apple"}
	usbArr[1] = Phone{"Google"}
	usbArr[2] = Camera{"Nikon"}

	//沒有賦值都是nil
	fmt.Println(usbArr)

}
