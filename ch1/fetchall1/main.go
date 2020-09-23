// Exercise 1.10: Find a web site that produces large amounts of data.
// Investigate caching by running fetchall twice in succession to see whether
// the reported time changes much. Do you get the same content each time?
// Modify fetchallto print its output to a file so it can be examined.

// See page 18.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
    file, err := os.Create("data.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "fetchall: %v\n", err)
        os.Exit(1)
    }

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
        fmt.Fprintf(file, "%v\n", <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

