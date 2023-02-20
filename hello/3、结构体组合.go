package main

import (
	"fmt"
)

// 结构体
type Person struct {
	Name string
}

type Saiyan1 struct {
	*Person
	Power int
}

func main() {
	// 使用
	use := &Saiyan1{
		Person: &Person{"test"},
		Power:  1000,
	}
	use.introduce()

	use1 := &Saiyan1{&Person{"myTest"}, 2000}
	myIntroduce(use1)

}

func (p *Person) introduce() {
	fmt.Println(p.Name)
}

func myIntroduce(s *Saiyan1) {
	fmt.Println(s.Person.Name)
}
