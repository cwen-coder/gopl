package main

import "fmt"

func delDup(s []string) []string {
	if len(s) == 0 {
		return s
	}
	d := s[0]
	out := s[:1]
	for i := 1; i < len(s); i++ {
		if d != s[i] {
			out = append(out, s[i])
			d = s[i]
		}
	}
	return out
}

func main() {
	x := []string{"aa", "aa", "bb", "cc", "bb"}
	fmt.Println(delDup(x))
}
