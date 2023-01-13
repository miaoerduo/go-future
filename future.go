package future

import (
	"sync"
	"time"
)

type Promise[T any] struct {
	data    T
	isReady bool
	mutex   sync.Mutex
	wg      sync.WaitGroup
}

type Future[T any] struct {
	p *Promise[T]
}

type FutureStatus int32

const (
	Ready   FutureStatus = 0
	Timeout FutureStatus = 1
)

func (s FutureStatus) ToString() string {
	switch s {
	case Ready:
		return "Ready"
	case Timeout:
		return "Timeout"
	default:
		return "Unknown"
	}
}

func NewPromise[T any]() *Promise[T] {
	p := &Promise[T]{}
	p.wg.Add(1)
	return p
}

func (p *Promise[T]) IsReady() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.isReady
}

func (p *Promise[T]) SetValue(data T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.isReady {
		panic("cannot set value more than once")
	}
	p.isReady = true
	p.data = data
	p.wg.Done()
}

func (p *Promise[T]) GetFuture() *Future[T] {
	return &Future[T]{p: p}
}

func (f *Future[T]) Get() T {
	f.Wait()
	return f.p.data
}

func (f *Future[T]) Wait() {
	f.p.wg.Wait()
}

func (f *Future[T]) WaitFor(timeoutDuration time.Duration) FutureStatus {
	c := make(chan struct{})
	go func() {
		defer close(c)
		f.p.wg.Wait()
	}()

	select {
	case <-c:
		return Ready
	case <-time.After(timeoutDuration):
		return Timeout
	}
}

func (f *Future[T]) WaitUntil(timeoutTime time.Time) FutureStatus {
	c := make(chan struct{})
	go func() {
		defer close(c)
		f.p.wg.Wait()
	}()

	select {
	case <-c:
		return Ready
	case <-time.After(time.Until(timeoutTime)):
		return Timeout
	}
}
