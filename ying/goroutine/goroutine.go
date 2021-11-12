package main

import (
	"fmt"
	"sync"
)

/**
go 并发的单位goroutine
*/
var wg sync.WaitGroup
var a int64 = 0

func add() {
	for i := 0; i < 5000; i++ {
		a += 1
	}
	wg.Done()
}

func main() {
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
