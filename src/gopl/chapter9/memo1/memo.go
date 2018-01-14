package memo

import (
    "sync"
)

// package memo provides a concurrency unsafe memoization of a function of type Func.


// A memo caches the result of calling a Func
type Memo struct {
    f Func
    mu sync.Mutex
    cache map[string]*entry
}

type entry struct {
    res result
    ready chan struct{} // closed when res is ready
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err error
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    e := memo.cache[key]
    if e == nil {
        // This is the first request for this key
        // This goroutine becomes responsible for computing
        // the value and broadcasting the ready condition
        e = &entry{ready: make(chan struct{})}
        memo.cache[key] = e
        memo.mu.Unlock()
        e.res.value, e.res.err = memo.f(key)
        close(e.ready)
    } else {
        // This is a repeat request for this key
        memo.mu.Unlock()
        <- e.ready // wait for ready condition
    }
    return e.res.value, e.res.err
}

