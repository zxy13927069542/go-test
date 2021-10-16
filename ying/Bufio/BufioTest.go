package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
使用Bufio读取文件
*/

func main() {

	Read()
	Write()

}

func Read() {
	//打开文件
	file, err := os.Open("./AAA.txt")
	if err != nil {
		fmt.Println("打开文件失败！错误原因为：", err)
		return
	}
	//使用Bufio读取文件
	fileReader := bufio.NewReader(file)
	for {
		//读取'\n'之前的数据并包括'\n'
		line, err := fileReader.ReadString('\n')
		//如果遇到err == io.EOF则返回错误并返回在遇到错误前的数据
		//当读到最后一行会报err == io.EOF，需要在此打印line
		if err == io.EOF {
			if len(line) == 0 {
				fmt.Println("ERROR！文件已经读完，无须再读！", err)
				break
			} else {
				fmt.Println(line)
				break
			}
		}
		//当且仅当找不到截止符号时会err != nil
		if err != nil {
			fmt.Println("ERROR! 找不到截止符号！", err)
			return
		}
		fmt.Println(line)
	}
	file.Close()
}

func Write() {
	file, err := os.OpenFile("./AAA.txt", os.O_WRONLY|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("ERROR! 打开文件错误！错误为：", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//先将数据写到缓冲，再一次性写入文件
	for i := 0; i < 10; i++ {
		writer.Write([]byte("bufiotest\n"))
	}
	writer.Flush()

}
