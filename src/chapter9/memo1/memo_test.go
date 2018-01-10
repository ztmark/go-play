package memo_test

import (
    "chapter9/memotest"
    "testing"
    "chapter9/memo1"
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