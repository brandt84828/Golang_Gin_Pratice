//這個可以過 雖然有重複的Test01()但方法簽章相同
package main

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

type CInterface interface {
	AInterface
	BInterface
}

func main() {

}


//這個會報錯 在CInterface有兩個重複的Test01()且方法簽章不相同
package main

type AInterface interface {
	Test01(a int)
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

type CInterface interface {
	AInterface
	BInterface
}

func main() {

}