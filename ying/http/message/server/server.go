package main

import (
	log "github.com/pion/ion-log"
	"io"
	"net/http"
	"os"
)

//
// Establish
//  @Description: 接收/establish
//
func Establish(w http.ResponseWriter, r *http.Request) {
	log.Infof("Receive /establish Message!")

	//  设置响应消息头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "Keep-Alive")

	//  设置响应消息体
	json := `{"version":"1.0","key":"95E4153B-F61F-43b3-917A-97C6950EE4B6"}`
	w.Write([]byte(json))
}

//
// HeartBeat
//  @Description: 接收/heartbeat
//
func HeartBeat(w http.ResponseWriter, r *http.Request) {
	log.Infof("Receive /heartbeat Message!")

	//  设置响应消息头
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Connection", "Keep-Alive")

}

func Push(w http.ResponseWriter, r *http.Request) {
	log.Infof("Receive /push Message!")
	io.Copy(os.Stdout, r.Body)
	r.Body.Close()

	w.Write([]byte("push success!"))

}

func Root(w http.ResponseWriter, r *http.Request) {
	log.Infof("Receive / Message!")
	io.Copy(os.Stdout, r.Body)
	r.Body.Close()

	w.Write([]byte("push success!"))

}

func Pulish(w http.ResponseWriter, r *http.Request) {
	log.Infof("Receive /publish Message!")

	w.Write([]byte("publish success!"))

}

func main() {
	//http.HandleFunc("/establish", Establish)
	//http.HandleFunc("/heartbeat", HeartBeat)
	//http.HandleFunc("/", Root)
	http.HandleFunc("/push", Push)

	//http.HandleFunc("/publish", Pulish)

	//for {
	//  如果设置ip的话会只监听目标主机的端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Errorf("%v", err)
	}
	//}

}
