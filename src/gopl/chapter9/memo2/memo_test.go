package memo

import (
    "chapter9/memotest"
    "testing"
    "chapter9/memo2"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Concurrent(t, m)
}

