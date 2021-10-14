package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
go 并发的单位goroutine
*/
var wg sync.WaitGroup
var a int64 = 0

func Hello_Goroutine() {
	//如果mian函数未阻塞，则此行尚未执行mian函数便已结束
	//mian函数结束，由其创建的goroutine也会结束
	//因此需要在mian函数中添加阻塞语句保证此函数能够执行完成
	fmt.Println("这里是goroutine!")
}

func Hello(i int) {
	//函数执行完线程池减一
	fmt.Println("这里是goroutine!", i)
	wg.Done()
}

//单向发送通道
func Send(chan1 chan<- string) {
	//往通道中发送数据
	chan1 <- "第一个数据" //可以使用len()获取通道中的元素个数，使用cap()获取通道的容量
	chan1 <- "第二个数据"
	chan1 <- "第三个数据"
	close(chan1) //即使通道关闭了也可以接收值
	wg.Done()
}

//单向接收通道
func Receive(chan1 <-chan string) {
	//使用for range遍历通道数据，当通道关闭时会自动退出for range,当通道中没值且通道未关闭时会阻塞
	for i := range chan1 {
		fmt.Println(i)
	}
	wg.Done()
}

func add() {
	for i := 0; i < 5000; i++ {
		a += 1
	}
	wg.Done()
}

func main() {
	//新建一个goroutine,该goroutine会与mian函数并发执行
	go Hello_Goroutine()
	//会比Hello_Goroutine()中的打印语句先执行，因为goroutine的创建需要时间
	fmt.Println("这里是mian函数！")
	//阻塞1s等待Hello_Goroutine()执行完成
	time.Sleep(time.Second)

	//runtime.GOMAXPROCS()  GOMAXPROCS(n int) int	设置程序并发时占用的CPU核心数
	//runtime.NumCPU()	NumCPU() int	获取当前程序所能使用的最大CPU核心数，在程序开始执行时便已确定
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(3)
	for i := 0; i < 33; i++ {
		wg.Add(1)
		go Hello(i)
	}
	//等待其他goroutine执行完
	wg.Wait()

	//通道声明
	//var chan1 chan interface{}
	//fmt.Println(chan1)	//channel通道为引用类型，零值为nil
	//通道初始化,第一个为通道类型，第二个为通道缓冲大小，不填缓冲则无缓冲,如果无缓冲，当前线程会阻塞直到通道中数据被接收
	//chan1 = make(chan interface{})
	//上述两步相当于下面一步
	chan1 := make(chan string, 3)
	fmt.Println(chan1)

	//发送数据
	wg.Add(1)
	go Send(chan1)

	//接收数据
	wg.Add(1)
	go Receive(chan1)

	wg.Wait()

	//select
	chan2 := make(chan int, 1)
	for i := 0; i < 10; i++ {
		//select每次仅执行一个case，如果有多个case满足，仅挑选出其中一个
		select {
		case data := <-chan2:
			fmt.Println(data)
		case chan2 <- i:
			fmt.Println("sent success")
		default:
			fmt.Println("无case执行")
		}
	}

	//两个线程竞争临界资源
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(a) //预计10000，结果9602

}
