//2,|gofmt
package main

import (
	"fmt"
	"os"
	"image"
	"github.com/9nut/iascii"
	_ "image/jpeg"
	_ "image/png"
)

func img2txt(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		return
	}
	defer f.Close()
	pic, _, err := image.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
		return
	}
	iascii.Encode(os.Stdout, pic)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file [...]\n", os.Args[0])
		os.Exit(1)
	}
	for _, fname := range os.Args[1:] {
		img2txt(fname)
	}
	os.Exit(0)
}
