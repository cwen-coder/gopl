package spaces

import (
	"unicode"
)

func trimSpaces(b []byte) []byte {
	runes := []rune(string(b))
	out := runes[:0]
	var isSpace bool
	for _, v := range runes {
		if unicode.IsSpace(v) && isSpace {
			continue
		} else if unicode.IsSpace(v) {
			out = append(out, ' ')
			isSpace = true
		} else {
			out = append(out, v)
			isSpace = false
		}
	}
	return []byte(string(out))
}
