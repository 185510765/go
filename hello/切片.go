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

}
