package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, des string) {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 777)
	if err != nil {
		fmt.Println("ERROR! 打开源文件失败！ 失败原因为：", err)
		return
	}
	//用于缓存源文件中的数据
	var buff []byte
	//读取源文件
	for {
		_, err := srcFile.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR! 读取源文件失败！ 失败原因为：", err)
			return
		}
	}
	srcFile.Close()

	//打开目标文件，如果目标文件不存在则创建，写文件为覆盖写
	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_WRONLY, 777)
	_, err = desFile.Write(buff)
	if err != nil {
		fmt.Println("ERROR! 写目标文件失败！ 失败原因为：", err)
		return
	}
	desFile.Close()

}
