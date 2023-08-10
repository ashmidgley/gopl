package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Please include filenames as arguments.")
		return
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			text := input.Text()
			counts[text]++
			if !contains(locations[text], file) {
				locations[text] = append(locations[text], file)
			}
		}
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, locations[line])
		}
	}
}

func contains(arr []string, file string) bool {
	for _, value := range arr {
		if value == file {
			return true
		}
	}
	return false
}
