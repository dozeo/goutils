package goutils

import (
	"time"
)

type Bench struct {
	start int64
	count int64
}

func (b *Bench) Start() {
	b.start = time.Now().UnixNano()
}

func (b *Bench) Stop() int64 {
	tmp := time.Now().UnixNano() - b.start
	b.count += tmp
	return tmp
}
func (b *Bench) Sum() int64 {
	return b.count
}
