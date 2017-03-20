package pool

/*
使用有缓冲的通道实现资源池，来管理可以在任意数量的goroutine之间共享及独立使用的资源。
 */

import (
	"errors"
	"sync"
	"io"
	"log"
)

type Pool struct {
	m  sync.Mutex
	resources chan io.Closer
	factory  func() (io.Closer, error)
	closed bool
}

var ErrPoolClosed  = errors.New("Pool has been closed.")

// 创建一个池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0{
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		factory:fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool)Acquire() (io.Closer, error) {
	select {
	case r, ok := <- p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok{
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release将一个使用后的资源放回池里
func (p *Pool)Release(r io.Closer)  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool)Close()  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources{
		r.Close()
	}
}