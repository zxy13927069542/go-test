package main

import (
	"fmt"
	log "github.com/pion/ion-log"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	client *http.Client
	once   sync.Once
)

func CreateHTTPClient() *http.Client {
	// 使用单例创建client
	once.Do(func() {
		client = &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        100,              // 最大连接数,默认0无穷大
				MaxIdleConnsPerHost: 100,              // 对每个host的最大连接数量(MaxIdleConnsPerHost<=MaxIdleConns)
				IdleConnTimeout:     90 * time.Second, // 多长时间未使用自动关闭连接
			},
		}
	})

	return client
}

func main() {
	client := CreateHTTPClient()

	// 测试/establish
	err := Establish(client)
	if err != nil {
		for err != nil {
			err = Establish(client)
		}
	}

	HeartBeat(client)

	//Push(client)

}

func Establish(client *http.Client) error {
	log.Infof("sending establish message...")
	req, err := http.NewRequest(http.MethodGet, "http://10.7.3.197:9999/establish", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "Keep-Alive")
	if err != nil {
		log.Errorf("new request failed! reason: %v", err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("request failed! reason: %v", err)
		return err
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	res.Body.Close()

	log.Infof("send establish message success!")
	return nil
}

func HeartBeat(client *http.Client) {
	for i := 0; i < 50; i++ {
		req, err := http.NewRequest(http.MethodGet, "http://10.7.3.197:9999/heartbeat", nil)
		req.Header.Set("Connection", "Keep-Alive")
		if err != nil {
		}
		res, err := client.Do(req)
		if err != nil {
		}

		bytes, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(bytes))

		res.Body.Close()
		time.Sleep(time.Second * 2)
	}

}

func Push(client *http.Client) error {
	req, err := http.NewRequest(http.MethodGet, "http://10.7.3.245:9999/push", nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Errorf("new request failed! reason: %v", err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("request failed! reason: %v", err)
		return err
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	res.Body.Close()
	return nil
}
