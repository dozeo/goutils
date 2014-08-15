package goutils

import (
	"math"
	"runtime"
)

type concurrencyLimit struct {
	lock chan (bool)
	size int
}

func NewConcurrencyLimit(limit int) concurrencyLimit {
	ob := concurrencyLimit{}
	ob.lock = make(chan bool, limit)
	ob.size = limit
	return ob
}

func NewConcurrencyLimitCPU() concurrencyLimit {
	return NewConcurrencyLimit(runtime.NumCPU())
}

func NewConcurrencyLimitCPUFactore(f float64) concurrencyLimit {
	t := float64(runtime.NumCPU()) * f
	t2 := math.Ceil(t)
	return NewConcurrencyLimit(int(t2))
}

func NewConcurrencyLimitOne() concurrencyLimit {
	return NewConcurrencyLimit(1)
}

func (l *concurrencyLimit) Use() {
	l.lock <- true
}

func (l *concurrencyLimit) Free() {
	<-l.lock
}

func (l *concurrencyLimit) Len() int {
	return len(l.lock)
}

func (l *concurrencyLimit) Size() int {
	return l.size
}

// usage: defer x.Limit()()
//                     ^ ^ double ()
func (l *concurrencyLimit) Limit() func() {
	l.Use()
	return func() {
		l.Free()
	}
}
