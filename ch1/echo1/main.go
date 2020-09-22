// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of 
// the command that invoked it.

// See page 8.

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args[1:])
    fmt.Println(os.Args[0])
}
