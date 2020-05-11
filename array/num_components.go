package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

func main() {
	f()
}

func f() {
	var out io.Writer
	out = bytes.NewBuffer([]byte{})
	fmt.Printf("%T %p\n", out, out)

	out = nil
	fmt.Printf("%T %p\n", out, out)
}
