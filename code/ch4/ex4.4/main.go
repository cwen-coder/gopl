package main

import "fmt"

func rotate(s []int, i int) {
	if len(s) == 0 {
		return
	}
	tmp := make([]int, len(s))
	copy(tmp, s)
	copy(s, s[i:])
	copy(s[len(s)-i:], tmp[:i])
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}
