package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	//	"io/ioutil"
	"os"
	"strings"
)

var bit = flag.Int("bit", 256, "bit of SHA. Default: 256.")

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	b := []byte(strings.Join(flag.Args(), " "))

	switch *bit {
	case 256:
		fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256(b))

	case 384:
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum384(b))
	case 512:
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512(b))
	default:
		fmt.Fprint(os.Stderr, "unsupported bit")
	}
}
