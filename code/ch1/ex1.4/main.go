package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string](map[string]int))
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, fileNames := range counts {
		lineCount := 0
		for _, filecount := range fileNames {
			lineCount += filecount
		}
		if lineCount > 1 {
			fmt.Printf("%d\t%s (", lineCount, line)
			sep := ""
			for name, _ := range fileNames {
				fmt.Printf("%s%s", sep, name)
				sep = " "
			}
			fmt.Println(")")
		}
	}
}

func countLines(f *os.File, name string, counts map[string](map[string]int)) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][name]++
	}
}
