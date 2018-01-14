package main

import (
    "io"
    "bytes"
    "log"
    "fmt"
)

// CustomLimitedReader must satisfy io.Closer
var _ io.Closer = (CustomLimitedReader)(nil)
type CustomLimitedReader struct {
    reader io.Reader
    N int64
}

func (r CustomLimitedReader) Read(p []byte) (n int, err error) {
    if r.N <= 0 {
        return 0, io.EOF
    }
    if int64(len(p)) > r.N {
        p = p[:r.N]
    }
    n, err = r.reader.Read(p)
    r.N -= int64(n)
    return n, err
}

func CustomLimitReader(r io.Reader, n int) io.Reader {
    return CustomLimitedReader{r, int64(n)}
}

func (r CustomLimitedReader) Close() error {
    return nil
}

func main() {
    var buf = bytes.NewBufferString("Hello")
    r := CustomLimitReader(buf, 3)
    data := [10]byte{1,2,3,97,65,66,67,68,69,70}
    n, err := r.Read(data[:])
    if err != nil {
        log.Println(err)
    }
    fmt.Println(n)
    fmt.Println(string(data[:]))
}
