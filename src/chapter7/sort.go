package main

import (
    "sort"
    "fmt"
    "time"
    "text/tabwriter"
    "os"
)

type StringSlice []string

func (p StringSlice) Len() int {
    return len(p)
}

func (p StringSlice) Less(i, j int) bool {
    return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

//====

type Track struct {
    Title string
    Artist string
    Album string
    Year int
    Length time.Duration
}

var tracks = []*Track{
    {"Go", "delilah", "From the roots UP", 2012, leng("3m38s")},
    {"Go", "Moby", "moby", 1992, leng("2m27s")},
    {"Go Ahead", "ALicia Keys", "As I am", 2007, leng("4m27s")},
    {"Ready 2 Go", "Martin Solveig", "Smash", 2011, leng("4m24s")},
}

func leng(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}

func printTrack(tracks []*Track) {
    const format = "%v\t%v\t%v\t%v\t%v\t\n"
    tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
    fmt.Fprintf(tw, format, "Title", "Artist", "ALbum", "Year", "Length")
    fmt.Fprintf(tw, format, "------", "------", "------", "------", "------")
    for _, t := range tracks {
        fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
    }
    tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int {
    return len(x)
}

func (x byArtist) Less(i, j int) bool {
    return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
    x[i], x[j] = x[j], x[i]
}

type customSort struct {
    t []*Track
    less func(x, y *Track) bool
}

func (x customSort) Len() int {
    return len(x.t)
}

func (x customSort) Less(i, j int) bool {
    return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
    x.t[i], x.t[j] = x.t[j], x.t[i]
}

func main() {
    names := []string{"jim", "mark", "tom", "frank"}
    fmt.Println(names)
    sort.Sort(StringSlice(names))
    //sort.Strings(names)
    fmt.Println(names)

    fmt.Println("=======")
    printTrack(tracks)
    fmt.Println("=======")
    sort.Sort(byArtist(tracks))
    printTrack(tracks)
    fmt.Println("=======")
    sort.Sort(sort.Reverse(byArtist(tracks)))
    printTrack(tracks)

    fmt.Println("=======")
    sort.Sort(customSort{tracks, func(x, y *Track) bool {
        if x.Title != y.Title {
            return x.Title < y.Title
        }
        if x.Year != y.Year {
            return x.Year < x.Year
        }
        if x.Length != y.Length {
            return x.Length < y.Length
        }
        return false
    }})
    printTrack(tracks)
}
