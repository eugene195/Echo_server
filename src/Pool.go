package main
import (
	"net"
)

type Pool struct {
	size  int
	tasks chan net.Conn
	kill  chan struct{}
}

func NewPool(size int) *Pool {
	pool := &Pool{
		tasks: make(chan net.Conn, 128),
		kill:  make(chan struct{}),
	}
	pool.Resize(size)
	return pool
}


func (p *Pool) Resize(n int) {
	for p.size < n {
		p.size++
		go p.Start()
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

func (p *Pool) Close() {
	close(p.tasks)
}


func (p *Pool) Exec(task net.Conn) {
	p.tasks <- task
}