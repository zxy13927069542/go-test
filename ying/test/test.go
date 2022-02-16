package main

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"
)

var HTTPTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 连接超时时间
		KeepAlive: 60 * time.Second, // 保持长连接的时间
	}).DialContext, // 设置连接的参数
	MaxIdleConns:          500,              // 最大空闲连接
	IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
	ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
	MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
}

func main() {
	text := `POST /push HTTP/1.1
Host: localhost
Content-Type: Content-Type: application/json
Content-Length: 426

{"topic":"CYGBase:RTDB-CYGBase:modify","body":"{\"name\":\"localsync\",\"node\":\".......................................-mingw-56\",\"topic\":\"CYGBase:RTDB-CYGBase:modify\",\"sn\":\"3F31D37D-9B5A-43c0-B6EA-76DC093BEED5\",\"data\":\"[{\\\"key\\\":\\\"CYGDW:Hash:GK:00505629ACBF:00505629ACBF\\\",\\\"tvModifys\\\":[{\\\"t\\\":\\\"online\\\",\\\"sv\\\":\\\"0\\\",\\\"dv\\\":\\\"0\\\",\\\"time\\\":\\\"1634527813090\\\"}]}]\"}"}`
	compile, _ := regexp.Compile(`\{.*}`)
	find := compile.Find([]byte(text))

	fmt.Println(string(find))
}
