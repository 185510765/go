package main

import (
	"fmt"
)

func main() {
	scores := [4]int{55, 66, 32, 88}
	fmt.Println(scores)
	fmt.Println(len(scores))

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	for _, value := range scores {
		// fmt.Println(index)
		fmt.Println(value)
	}
}
