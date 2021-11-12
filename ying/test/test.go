package main

import (
	"fmt"
	"sync"
)

/**
使用并发模拟接力比赛
*/

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ch := make(chan int)
	go Runner(ch)

	//比赛开始
	ch <- 1
	wg.Wait()
}

func Runner(ch chan int) {
	defer wg.Done()
	runner := <-ch
	fmt.Printf("运动员%d号接棒\n", runner)
	fmt.Printf("运动员%d号跑起来了\n", runner)
	if runner < 4 {
		newRunner := runner + 1
		fmt.Printf("运动员%d号已就位\n", newRunner)
		wg.Add(1)
		go Runner(ch)

		fmt.Printf("运动员%d号开始传棒\n", runner)
		ch <- newRunner
	} else {
		fmt.Printf("运动员%d号冲过终点\n", runner)
		fmt.Println("比赛结束！")
		close(ch)
	}

}
