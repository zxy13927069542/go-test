package main

import (
	"fmt"
	"go-test/ying/error"
	"io"
	"os"
)

/**
使用file.Read()完整读取文件
以及err = io.EOF的复现测试
*/

func main() {

	//调用写文件操作
	//Write()
	//调用文件读操作
	Read()
	//调用文件复制函数
	//CopyFile("ying/error/ErrorMsg.go", "ErrorMsg.txt")

}

func Read() {
	file, err := os.Open("ErrorMsg.txt")
	if err != nil {
		error.ErrorMsg("打开文件失败！", err)
	}
	defer file.Close()
	//接受文件内容的缓冲切片
	buff := make([]byte, 128)
	var contend []byte
	for {
		n, err := file.Read(buff)
		if err == io.
			EOF {
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

	////文件内部读取存在一个光标，指示当前读取到文件的哪里，如果文件读完后再次进行读取则会报错，error = io.EOF
	//n,error := file.Read(buff)
	//fmt.Println(string(buff[:n]))
	//_,error = file.Read(buff)
	//if error == io.EOF {
	//	fmt.Println("文件读完了！")
	//	return
}

// Write 在AAA.txt文件写10行内容，写方式为覆盖写，写之前会清空文件内容
func Write() {
	//使用权限打开文件，使文件拥有读写相关的权限
	//os.O_TRUNC 清空
	//os.O_WRONLY 写权限
	file, err := os.OpenFile("./AAA.txt", os.O_WRONLY|os.O_TRUNC, 777)
	if err != nil {
		error.ErrorMsg("打开文件失败！", err)

	}
	defer file.Close()
	//默认直接从第一个字节开始写，如果文件有内容会被替换
	//os.O_APPEND 追加在已有内容之后

	for i := 0; i < 10; i++ {
		_, err := file.WriteString("this is FileReadingTest.Write!\n")
		if err != nil {
			error.ErrorMsg("写文件失败！", err)
		}
	}
}
