package worker
/*
使用无缓冲的通道来创建一个goroutine池，
这些goroutine执行并控制一组工作，让其并发执行。
 */

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

// New 一个新的工作池
func New(maxGoroutines int) *Pool {
	p:=Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i:=0;i<maxGoroutines;i++ {
		go func() {
			for w:=range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool)Run(w Worker)  {
	p.work <- w
}
// Shutdown 等待所有goroutine停止工作
func (p *Pool)Shutdown()  {
	close(p.work)
	p.wg.Wait()
}