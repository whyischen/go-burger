package main

import (
	"testing"
)

func TestSMap(t *testing.T) {
	err := sMap()
	if err != nil {
		t.Error(err)
	}
}

// 互斥锁
// mutex := sync.Mutex{}
// mutex.Lock()
// 读写锁
// rwMutex := sync.RWMutex{}
// 读锁 rwMutex.RLock()
// 写锁 rwMutex.Lock()
func TestLock(t *testing.T) {
}
