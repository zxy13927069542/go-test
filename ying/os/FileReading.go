package main

import (
	"fmt"
	"io"
	"os"
)

/**
使用file.Read()完整读取文件
以及err = io.EOF的复现测试
*/

func main() {

	//调用写文件操作
	Write()
	//调用文件复制函数
	CopyFile("./AAA.txt", "./BBB.txt")

}

func Read() {
	file, err := os.Open("./AAA.txt")
	if err != nil {
		fmt.Println("文件打开失败，错误原因为：", err)
	}
	defer file.Close()
	//接受文件内容的缓冲切片
	buff := make([]byte, 128)
	var contend []byte
	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			fmt.Println("文件已经读完了，无须再读！")
			break
		}
		if err != nil {
			fmt.Println("文件读取失败，错误为：", err)
			return
		}
		contend = append(contend, buff[:n]...)
	}
	fmt.Println(string(contend))

	////文件内部读取存在一个光标，指示当前读取到文件的哪里，如果文件读完后再次进行读取则会报错，err = io.EOF
	//n,err := file.Read(buff)
	//fmt.Println(string(buff[:n]))
	//_,err = file.Read(buff)
	//if err == io.EOF {
	//	fmt.Println("文件读完了！")
	//	return
}

func Write() {
	//使用权限打开文件，使文件拥有读写相关的权限
	file, err := os.OpenFile("./AAA.txt", os.O_WRONLY|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("ERROR! 打开文件错误！错误类型为：", err)
		return
	}
	defer file.Close()
	//默认直接从第一个字节开始写，如果文件有内容会被替换
	//os.O_APPEND 追加在已有内容之后
	file.Write([]byte("我是大人莫"))
}
