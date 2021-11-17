package main

import (
	"fmt"
	"io"
	"os"
)

/**
复制文件
*/
func CopyFile(src, des string) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 777)
	if err != nil {
		fmt.Println("ERROR! 打开源文件失败！ 失败原因为：", err)
		return
	}
	//用于缓存源文件中的数据
	buff := make([]byte, 128)
	//var buff []byte	//无法使用nil的切片来接收
	var contend []byte
	//读取源文件
	for {
		n, err := srcFile.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR! 读取源文件失败！ 失败原因为：", err)
			return
		}
		contend = append(contend, buff[:n]...) //append()会自动扩容，所以可以用nil
	}
	srcFile.Close()

	//打开目标文件，如果目标文件不存在则创建，写文件为覆盖写
	//os.O_CREATE
	//注意这里perm要写0666，如果写777在文件存在的情况下会返回nil
	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	//_, err = desFile.Write(contend)
	_, err = desFile.Write(contend)
	if err != nil {
		fmt.Println("ERROR! 写目标文件失败！ 失败原因为：", err)
		return
	}
	desFile.Close()

}
