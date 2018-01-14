package main

import (
    "os"
    "io/ioutil"
    "fmt"
    "path/filepath"
    "flag"
    "time"
    "sync"
)

// prevent too much goroutine
var sema = make(chan struct{}, 20)

func walkDirectory(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
    defer n.Done()
    for _, entry := range ls(dir) {
        if entry.IsDir() {
            n.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDirectory(subdir, n, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

func ls(dir string) []os.FileInfo {
    sema <- struct{}{} // acquire token
    defer func() { <-sema }() // release token
    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1:%v\n", err)
        return nil
    }
    return entries
}

var v = flag.Bool("v", false, "show verbose progress message")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }
    // Traverse each root of the file tree in parallel
    filesSizes := make(chan int64)
    var n sync.WaitGroup
    for _, root := range roots {
        n.Add(1)
        go walkDirectory(root, &n, filesSizes)
    }
    go func() {
        n.Wait()
        close(filesSizes)
    }()

    // Print the result period
    var tick <-chan time.Time
    if *v {
        tick = time.Tick(500 * time.Millisecond)
    }
    var nfiles, nbytes int64
loop:
    for {
        select {
        case size, ok := <-filesSizes:
            if !ok {
                break loop
            }
            nfiles++
            nbytes += size
        case <-tick:
            printDU(nfiles, nbytes)
        }
    }
    printDU(nfiles, nbytes)
}

func printDU(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f KB\n", nfiles, float64(nbytes)/ 1e3)
}
