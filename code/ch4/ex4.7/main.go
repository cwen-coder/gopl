package main

import (
	"fmt"
)

func reverse(b []byte) {
	for i, j := 0, len(b)-3; i < j; i, j = i+3, j-3 {
		b[i], b[j] = b[j], b[i]
		b[i+1], b[j+1] = b[j+1], b[i+1]
		b[i+2], b[j+2] = b[j+2], b[i+2]
	}
}

func main() {
	s := []byte("你好世界")
	reverse(s)
	fmt.Println(string(s))
}
