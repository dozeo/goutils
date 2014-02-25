package goutils

import (
	"container/list"
	"strconv"
	"sync"
	"time"
)

var Sync Synchronized

type Synchronized struct {
	globalHashLock sync.Mutex
	locks          list.List
}

func (s *Synchronized) Lock(key interface{}) {
	wait := true
outer:
	for wait {
		s.globalHashLock.Lock()
		for e := s.locks.Front(); e != nil; e = e.Next() {
			if e.Value == key {
				s.globalHashLock.Unlock()
				time.Sleep(1 * time.Millisecond)
				continue outer
			}
		}
		wait = false
	}
	s.locks.PushFront(key)
	s.globalHashLock.Unlock()
}

func (s *Synchronized) Unlock(key interface{}) {
	removed := false
	s.globalHashLock.Lock()
	for e := s.locks.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			s.locks.Remove(e)
			removed = true
			break
		}
	}
	s.globalHashLock.Unlock()
	if removed == false {
		panic("unlock of unlocked syncronized. known locks: " + strconv.Itoa(s.locks.Len()))
	}
}

func (s *Synchronized) Call(lock interface{}, call func()) {
	s.Lock(lock)
	defer s.Unlock(lock)
	call()
}
