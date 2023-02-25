package main

import (
	"fmt"
	"os"
)

func main() {
	// 目录操作
	os.Mkdir("file", 0777)
	os.MkdirAll("file/file1/file2", 0777)
	// err := os.Remove("file/file1/file2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// os.RemoveAll("file")

	// 文件操作 写文件
	fileName := "file/file1/file2/create.txt"
	fout, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	for i := 0; i < 5; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	// 读文件
	fread, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fread.Close()
	buf := make([]byte, 2014)
	for {
		n, _ := fread.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
