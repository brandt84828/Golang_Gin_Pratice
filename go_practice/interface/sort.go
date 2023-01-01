package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.聲明Hero結構體
type Hero struct {
	Name string
	Age  int
}

// 2.聲明Hero結構體切片類型
type HeroSlice []Hero

// 3.實現Interface
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法就是決定使用什麼標準進行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//也可以修改成對Name進行排序
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	//更快速的方法(等同上面3句) hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//先定義切片
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("Hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//將hero append到heros切片
		heroes = append(heroes, hero)
	}

	//查看排序前順序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//調用sort.sort
	sort.Sort(heroes)

	fmt.Println("------------------------------------")
	//查看排序後順序
	for _, v := range heroes {
		fmt.Println(v)
	}

	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i=", i, "j=", j) //i=20 j=10

}
