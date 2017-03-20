package main

import (
	"os"
	"time"
	"errors"
	"os/signal"
	"log"
)

// runner 用于展示如何使用通道来监视程序的执行时间
// 如果执行时间过长，可以用来终止程序

// Runner在给定的超时时间内执行一组任务，
// 并且在操作系统发送中断信号时结束这些任务。


const timeout = 3 * time.Second

type Runner struct {
	// interrupt 从操作系统发送打断信号
	interrupt chan os.Signal
	// finish 通道报告处理任务已经完成
	finish chan error

	// timeout 报告处理任务已经超时
	timeout <- chan time.Time

	// tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在接收到系统发送的超时信号时返回
var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

// New_runner 返回一个新的Runner
func New_runner(d time.Duration)  *Runner {
	return &Runner{
		interrupt:make(chan os.Signal, 1),
		finish: make(chan error),
		timeout: time.After(d),
	}
}

// Add 将一个任务附加到Runner上。
// 这个任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner)Add(tasks ...func(int))  {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行，并监控通道
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的goroutine执行不同的任务
	go func() {
		r.finish <- r.run()
	}()
	select {
	case err:=<-r.finish:return err
	case <-r.timeout: return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner)run() error {
	for id, task:=range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner)gotInterrupt() bool {
	select {
	case <- r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process -Task %v.", id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}

func main() {
	log.Println("开始工作")
	r := New_runner(timeout)
	// 加入要执行任务并处理结果
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil{
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}
