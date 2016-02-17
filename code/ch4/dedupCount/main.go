package main

import (
	"bufio"
	"fmt"
	"os"
)

var m = make(map[string]int)

//func K(list []string) string {
//	return fmt.Sprintf("%q", list)
//}

func Add(list string) {
	m[list]++
}

func Count(list string) int {
	return m[list]
}

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		Add(line)
		fmt.Println(Count(line))
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
