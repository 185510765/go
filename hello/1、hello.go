package main

import (
	"fmt"
	"os"
	"time"
)

// 结构体
type Saiyan struct {
	Name  string
	Power int
}

func main() {
	// 变量 ******************************************************
	var power int = 9000
	fmt.Printf("It's over %d\n", power)

	// := ，它可以自动推断变量类型，但是这种声明运算符只能用于局部变量，不可用于全局变量
	power2 := getPower()
	fmt.Println(power2)

	name, power := "GO", 90002
	fmt.Println(name, power)

	// 导入包 ****************************************************
	if len(os.Args) != 2 {
		// os.Exit(1)
	}
	fmt.Println("It's over", os.Args[0])

	// 函数声明 ****************************************************
	_, exists := showFunc("张三")
	if exists == false {
		fmt.Println("仅仅取其中的一个返回值")
	}

	// 结构体 *************************************************
	// goku := Saiyan{
	// 	Name:  "GoKu",
	// 	Power: 9000,
	// }
	// goku1 := Saiyan{Name: "Goku"}
	// goku1.Power = 9000

	// 指针
	goku := &Saiyan{"power", 9000}
	super(goku)
	fmt.Println(goku.Name)
	fmt.Println(goku.Power)

	// 结构体上面的函数
	sai := Saiyan{"测试账号", 8808}
	fmt.Println(sai)
	sai.displaySaiyan()

	// new
	goku3 := &Saiyan{
		Name:  "张三",
		Power: 8800,
	}
	fmt.Println(goku3)

	// 赋值技巧 分组赋值*********************************************
	const (
		i = 100
		f = 3.1415
		s = "GO_"
	)

	var (
	// i1     int
	// pi     float32
	// prefix string
	)

	for i := 0; i < 20; i++ {
		if i == 3 {
			// break
			// continue
		}

		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

}

// 函数 ******************************************************************
// displaySaiyan() 方法将 Saiyan 做为接收器类型
func (s Saiyan) displaySaiyan() {
	fmt.Println(s)
}

func super(s *Saiyan) {
	s.Power += 10000
}

func getPower() int {
	return 9001
}

func showFunc(name string) (int, bool) {
	return 100, false
}

// 参数类型相同 简写
func add(a, b int) int {
	return a + b
}
