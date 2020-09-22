// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package, and Section 11.4 shows
// how to write benchmark tests for systematic performance evaluation.)

// See page 8.

package main

import (
    "fmt"
    "os"
    "time"
    "strings"
)

func main() {
    start := time.Now()
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
    fmt.Printf("%f elapsed\n", time.Since(start).Seconds())

    start = time.Now()
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Printf("%f elapsed\n", time.Since(start).Seconds())
}
