package comma

import (
	"strings"
)

func comma(s string) string {
	var sign string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = string(s[0])
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	var suffix string
	if dot > -1 {
		suffix = s[dot:]
		s = s[:dot]
	}
	n := len(s)
	if n <= 3 {
		return sign + s + suffix
	}
	return sign + comma(s[:n-3]) + "," + s[n-3:] + suffix
}
