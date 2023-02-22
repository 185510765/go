package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		// os.Exit(1)
	}

	err := process(0)
	if err != nil {
		fmt.Println(err)
	}

	n, err := strconv.Atoi(os.Args[0])
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

}

// 我们可以通过导入 errors 包然后使用它的 New 函数创建我们自己的错误：
func process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}

	return nil
}
