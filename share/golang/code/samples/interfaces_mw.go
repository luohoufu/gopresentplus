package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h := crc32.NewIEEE()
	mw := io.MultiWriter(h, os.Stdout)
	fmt.Fprintf(mw, "Gotta write 'em all\n")
	fmt.Printf("hash=%#x\n", h.Sum32())
}
