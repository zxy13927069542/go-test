package main

import (
	"fmt"
	"io/ioutil"
)

/**
使用ioutil.ReadFile()读取文件
go 1.16开始，该方法相当于os.ReadFile()
*/

func main() {

	Read()
	Write()

}

func Read() {
	contend, err := ioutil.ReadFile("./os/FileReading.go")
	if err != nil {
		fmt.Println("ERROR! 读取文件错误！错误为：", err)
		return
	}
	fmt.Println(string(contend))
}

/**
ioutil的WriteFile()会直接覆盖文件内容
相当于os.WriteFile(filename, data, perm)
OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, perm)
*/
func Write() {
	var buff []byte
	for i := 0; i < 10; i++ {
		buff = append(buff, []byte("ioutiltest\n")...)
	}
	err := ioutil.WriteFile("./AAA.txt", buff, 777)
	if err != nil {
		fmt.Println("ERROR! 写文件错误！错误为：", err)
		return
	}

}
