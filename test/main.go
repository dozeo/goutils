package main

import (
	"github.com/dozeo/synchronized"
	"runtime"
	"time"
)

func a(v int) {
	k++
	for b := 0; b < 10000; b++ {
		synchronized.Lock("a")
		synchronized.Unlock("a")
	}
	for b := 0; b < 10000; b++ {
		synchronized.Synchronized("lockname", func() {
		})
	}
	k--
}

var k int = 0

func main() {
	runtime.GOMAXPROCS(20)
	for i := 0; i < 1000; i++ {
		go a(i)
	}
	for k > 0 {
		time.Sleep(1 * time.Millisecond)
	}
}
