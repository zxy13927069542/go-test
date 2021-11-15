package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

/**
Runnerr在给点的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
*/
type Runner struct {
	//interrupt接收操作系统发送的信号
	interrupt chan os.Signal

	//complete接收处理任务完成信号
	complete chan error

	//timeout报告任务已经超时
	timeout <-chan time.Time

	//tasks代表一组顺序执行的任务
	tasks []func(int)
}

//创建超时错误，类似于java的exception
var ErrTimeOut = errors.New("received timeout")

//创建中断错误，类似于java的exception
var ErrInterrupt = errors.New("received interrupt")

var Complete = errors.New("mission complete")

//Runner构造函数
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), //接收系统发过来的中断信号，要设缓冲为1
		complete:  make(chan error),
		//func After(d Duration) <-chan Time
		//After会在经过时间段d后往返回的通道发送当前时间
		timeout: time.After(d),
	}
}

//新增任务
func (r *Runner) Add(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

//执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	//使用r.interrupt接收系统的中断信号os.Interrupt
	signal.Notify(r.interrupt, os.Interrupt)

	//并发执行任务
	go func() {
		r.complete <- r.run()
	}()

	//处理任务完成的信号
	select {
	//处理完成信号
	case err := <-r.complete:
		return err
		//处理超时信号
	case <-r.timeout:
		return ErrTimeOut
	}
}

/**
执行任务
*/
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return Complete
}

//判断是否有来自系统的中断信号，有返回true，无返回false
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func task(id int) {
	fmt.Println("这里是任务：", id)
	time.Sleep(time.Second)
}

func main() {
	var tasks []func(int)
	for i := 0; i < 10; i++ {
		tasks = append(tasks, task)
	}

	runner := New(time.Second * 30)
	runner.Add(tasks...)
	err := runner.Start()
	if err != nil {
		fmt.Println(err)
	}
}
