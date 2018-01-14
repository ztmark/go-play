package memo

import "sync"

// package memo provides a concurrency unsafe memoization of a function of type Func.


// A memo caches the result of calling a Func
type Memo struct {
    f Func
    mu sync.Mutex
    cache map[string]result
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err error
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    res, ok := memo.cache[key]
    memo.mu.Unlock()
    if !ok {
        res.value, res.err = memo.f(key)
        memo.mu.Lock()
        memo.cache[key] = res
        memo.mu.Unlock()
    }
    return res.value, res.err
}

