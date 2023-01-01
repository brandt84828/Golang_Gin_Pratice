package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("bool : %v,%v\n", index, x)
		case float32:
			fmt.Printf("float32 : %v,%v\n", index, x)
		case float64:
			fmt.Printf("float64 : %v,%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("int : %v,%v\n", index, x)
		case string:
			fmt.Printf("string : %v,%v\n", index, x)
		case Student:
			fmt.Printf("Student : %v,%v\n", index, x)
		case *Student:
			fmt.Printf("*Student : %v,%v\n", index, x)
		default:
			fmt.Printf("%v,%v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "Taiwan"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
