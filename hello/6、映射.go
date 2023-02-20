package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["张三"] = 88
	lookup["李四"] = 94
	fmt.Println(lookup)

	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)

	len := len(lookup)
	fmt.Println(len)

	delete(lookup, "张三")
	fmt.Println(lookup)
	delete(lookup, "李四")
	fmt.Println(lookup)

	// 初始大小
	myMap := make(map[string]int, 100)
	fmt.Println(myMap)

	// 映射结构体
	type Student struct {
		Name    string
		Sex     int
		Score   int
		Friends map[string]*Student
	}
	student := &Student{
		Name:    "王五",
		Sex:     1,
		Score:   89,
		Friends: make(map[string]*Student),
	}
	// student.Friends["krillin"] =
	fmt.Println(student)
}
