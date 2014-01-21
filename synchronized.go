package synchronized

import (
	"container/list"
	"sync"
	"time"
)

var globalHashLock sync.Mutex
var locks list.List

func Lock(key interface{}) {
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

func Unlock(key interface{}) {
	removed := false
	globalHashLock.Lock()
	for e := locks.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			locks.Remove(e)
			removed = true
			break
		}
	}
	globalHashLock.Unlock()
	if removed == false {
		panic("unlock of unlocked syncronized")
	}
}

func Synchronized(lock interface{}, call func()) {
	Lock(lock)
	defer Unlock(lock)
	call()
}
