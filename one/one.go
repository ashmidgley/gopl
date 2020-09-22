package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    three()

    // 1.1 
    fmt.Println(os.Args[0])

    // 1.2
    for i, arg := range os.Args[1:] {
        fmt.Printf("%d %s\n", i, arg)
    }

    // 1.3
    start := time.Now()
    one()
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

    start = time.Now()
    two()
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func one() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func two() {
    fmt.Println(strings.Join(os.Args[1:], ""))
}

func three() {
    fmt.Println(os.Args[1:])
}
