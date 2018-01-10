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

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    defer memo.mu.Unlock()
    res, ok := memo.cache[key]
    if !ok {
        res.value, res.err = memo.f(key)
        memo.cache[key] = res
    }
    return res.value, res.err
}

