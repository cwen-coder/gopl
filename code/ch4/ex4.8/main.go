package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	CHAR_IS_SPACE = iota
	CHAR_IS_SYMBOL
	CHAR_IS_MARK
	CHAR_IS_DIGIT
	CHAR_IS_PUNCT
	CHAR_IS_LETTER
	CHAR_IS_NUMBER
	CHAR_IS_CONTROL
	CHAR_IS_GRAPHIC
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if r == '0' {
			break
		}

		switch {
		case unicode.IsSpace(r):
			counts[CHAR_IS_SPACE]++
		case unicode.IsSymbol(r):
			counts[CHAR_IS_SYMBOL]++
		case unicode.IsMark(r):
			counts[CHAR_IS_MARK]++
		case unicode.IsDigit(r):
			counts[CHAR_IS_DIGIT]++
		case unicode.IsPunct(r):
			counts[CHAR_IS_PUNCT]++
		case unicode.IsLetter(r):
			counts[CHAR_IS_LETTER]++
		case unicode.IsNumber(r):
			counts[CHAR_IS_NUMBER]++
		case unicode.IsControl(r):
			counts[CHAR_IS_CONTROL]++
		case unicode.IsGraphic(r):
			counts[CHAR_IS_GRAPHIC]++
		}
		utflen[n]++
	}
	fmt.Printf("type\tcount\n")
	for c, n := range counts {
		var name string
		switch c {
		case CHAR_IS_SPACE:
			name = "space"
		case CHAR_IS_SYMBOL:
			name = "symbol"
		case CHAR_IS_MARK:
			name = "mark"
		case CHAR_IS_DIGIT:
			name = "digit"
		case CHAR_IS_PUNCT:
			name = "punct"
		case CHAR_IS_LETTER:
			name = "letter"
		case CHAR_IS_NUMBER:
			name = "number"
		case CHAR_IS_CONTROL:
			name = "control"
		case CHAR_IS_GRAPHIC:
			name = "graphic"
		}
		fmt.Printf("%q\t%d\n", name, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
