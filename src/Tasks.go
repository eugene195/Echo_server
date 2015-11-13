package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//type Task interface {
//	Execute()
//}
////
////type Pool struct {
////	mu    sync.Mutex
////	size  int
////	tasks chan Task
////	kill  chan struct{}
////	wg    sync.WaitGroup
////}
//
////func NewPool(size int) *Pool {
////	pool := &Pool{
////		tasks: make(chan Task, 128),
////		kill:  make(chan struct{}),
////	}
////	pool.Resize(size)
////	return pool
////}
////
////func (p *Pool) worker() {
////	defer p.wg.Done()
////	for {
////		select {
////		case task, ok := <-p.tasks:
////			if !ok {
////				return
////			}
////			task.Execute()
////		case <-p.kill:
////			return
////		}
////	}
////}
//
//
//
//
//
//
//func main() {
//	pool := NewPool(5)
//
//	pool.Exec(ExampleTask("foo"))
//	pool.Exec(ExampleTask("bar"))
//
//	pool.Resize(3)
//
//	pool.Resize(6)
//
//	for i := 0; i < 20; i++ {
//		pool.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
//	}
//
//	pool.Close()
//
//	pool.Wait()
//}