package resty_test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"testing"
)

//测试resty库

func TestResty(t *testing.T) {
	//创建一个客户端
	client := resty.New()
	//调用客户端创建一个请求
	url := "http://www.baidu.com"
	resp, err := client.R().Get(url)
	if err != nil {
		log.Fatalln("请求", url, "失败", err)
	}

	fmt.Println("StatusCode:", resp.StatusCode())
	str := string(resp.Body())
	fmt.Printf(str)
	// all, err := ioutil.ReadAll(response.RawResponse.Body)
	// response.RawResponse.Body.Close()
	// if err != nil {
	// 	log.Fatalln("读取response.RawResponse.Body失败", err)
	// }
	// fmt.Println("RawBody:", string(all))

}
