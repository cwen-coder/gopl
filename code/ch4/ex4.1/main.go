package main

import (
	"crypto/sha256"
	"fmt"
)

func diff(x, y [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		for j := uint(0); j < 8; j++ {
			fmt.Println(x[i]&byte(1<<j), y[i]&byte(1<<j))
			if x[i]&byte(1<<j) == y[i]&byte(1<<j) {
				count++
			}
		}
	}
	return count
}

func main() {
	x := sha256.Sum256([]byte("a"))
	y := sha256.Sum256([]byte("b"))
	fmt.Printf("%x\n%x\n%d\n", x, y, diff(x, y))
}
