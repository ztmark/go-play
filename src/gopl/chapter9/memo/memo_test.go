package memo

import (
    "chapter9/memotest"
    "testing"
    "chapter9/memo"
)

var httpGetBody = memotest.HTTPGetBody

func TestMemoSequential(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Sequential(t, m)
}

func TestMemoConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Concurrent(t, m)
}