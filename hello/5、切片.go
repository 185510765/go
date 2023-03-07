package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 3, 22, 58, 222}
	fmt.Println(scores)

	scores1 := make([]int, 0, 10)
	// scores[7] = 2222
	scores1 = append(scores1, 5)
	fmt.Println(scores1)

	scores3 := make([]int, 0, 10)
	scores3 = scores3[0:8]
	scores3[7] = 9033
	fmt.Println(scores3)

	// 切片长度和容量
	slice := make([]string, 2, 3)
	fmt.Println("长度:", len(slice), "容量：", cap(slice))

	newSlice := slice[1:2:3]
	fmt.Println(newSlice)

	// // copy
	// num := make([]int, 100)
	// for i := 0; i < 100; i++ {
	// 	num[i] = int(rand.Int31n(1000))
	// 	num.Println(num[i])
	// }
	// sort.Ints(num)

	// fmt.Println(num)
	// fmt.Println(sort)
}
