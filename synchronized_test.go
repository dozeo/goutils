package goutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"runtime"
	"testing"
	"time"
)

func a(v int) {
	k++
	for b := 0; b < 10000; b++ {
		Sync.Lock("a")
		Sync.Unlock("a")
	}
	for b := 0; b < 10000; b++ {
		Sync.Call("lockname", func() {
		})
	}
	k--
}

var k int = 0

func Test_Sync(t *testing.T) {
	Convey("generate sync", t, func() {
		runtime.GOMAXPROCS(20)
		for i := 0; i < 100; i++ {
			go a(i)
		}
		for k > 0 {
			time.Sleep(1 * time.Millisecond)
		}
	})
}
