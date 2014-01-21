package synchronized

import (
	"container/list"
	"sync"
	"time"
)

var globalHashLock sync.Mutex
var locks list.List

func Lock(key string) {
	wait := true
outer:
	for wait {
		globalHashLock.Lock()
		for e := locks.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				globalHashLock.Unlock()
				time.Sleep(1 * time.Millisecond)
				continue outer
			}
		}
		wait = false
	}
	locks.PushFront(key)
	globalHashLock.Unlock()
}

func Unlock(key string) {
	removed := false
	for e := locks.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			globalHashLock.Lock()
			locks.Remove(e)
			globalHashLock.Unlock()
			removed = true
			break
		}
	}
	if removed == false {
		panic("unlock of unlocked syncronized")
	}
}