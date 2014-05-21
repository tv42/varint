package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", prog)
	fmt.Fprintf(os.Stderr, "  %s <FILE\n", prog)
	flag.PrintDefaults()
}

func main() {
	prog := path.Base(prog)
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		n, err := binary.ReadVarint(r)
		if err == io.EOF && n != 0 {
			// this only happens if encoding/binary read some bytes,
			// but not enough for a full varint
			log.Fatal("truncated input")
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Println(n)
		if err != nil {
			log.Fatalf("cannot write to standard output: %s", err)
		}
	}
}
