package main

import (
	"fmt"
)

func main() {
	Int := [...]int{1, 2, 3, 4, 5, 6, 7}
	a := Int[1:3]
	b := Int[2:3]
	a[1] = 9
	fmt.Println(Int)
	fmt.Println(a)
	fmt.Println(b)
}
