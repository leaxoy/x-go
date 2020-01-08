package xsync

import (
	"sync"
)

type AnyMutex struct {
	any interface{}
	sync.RWMutex
}

func NewAnyMutex(any interface{}) *AnyMutex {
	return &AnyMutex{any: any}
}

func (a *AnyMutex) Load() interface{} {
	a.RLock()
	defer a.RUnlock()
	return a.any
}

func (a *AnyMutex) Store(any interface{}) {
	a.Lock()
	defer a.Unlock()
	a.any = any
}
