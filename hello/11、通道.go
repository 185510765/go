package main

import "fmt"

func main() {
	c := make(chan int)

}

type Worker struct {
	id int
}

func (w Worker) process(c chan int) {}

// 将通道传递给函数
func worker(c chan int) {
	fmt.Println(c)
}
