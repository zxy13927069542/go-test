package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Http(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Get失败", err)
		return
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll()失败", err)
		return
	}

	fmt.Println(string(all))

}
