package goutils

import (
	"math"
	"runtime"
)

type ConcurrencyLimit struct {
	lock chan (bool)
	size int
}

func NewConcurrencyLimit(limit int) ConcurrencyLimit {
	ob := ConcurrencyLimit{}
	ob.lock = make(chan bool, limit)
	ob.size = limit
	return ob
}

func NewConcurrencyLimitCPU() ConcurrencyLimit {
	return NewConcurrencyLimit(runtime.NumCPU())
}

func NewConcurrencyLimitCPUFactore(f float64) ConcurrencyLimit {
	t := float64(runtime.NumCPU()) * f
	t2 := math.Ceil(t)
	return NewConcurrencyLimit(int(t2))
}

func NewConcurrencyLimitOne() ConcurrencyLimit {
	return NewConcurrencyLimit(1)
}

func (l *ConcurrencyLimit) Use() {
	l.lock <- true
}

func (l *ConcurrencyLimit) Free() {
	<-l.lock
}

func (l *ConcurrencyLimit) Len() int {
	return len(l.lock)
}

func (l *ConcurrencyLimit) Size() int {
	return l.size
}

// usage: defer x.Limit()()
//                     ^ ^ double ()
func (l *ConcurrencyLimit) Limit() func() {
	l.Use()
	return func() {
		l.Free()
	}
}
