package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", prog)
	fmt.Fprintf(os.Stderr, "  %s NUMBER..\n", prog)
	flag.PrintDefaults()
}

func main() {
	prog := path.Base(prog)
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
		os.Exit(1)
	}

	var buf [binary.MaxVarintLen64]byte
	for _, s := range flag.Args() {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("cannot parse number: %s", err)
		}

		n := binary.PutVarint(buf[:], num)
		_, err = os.Stdout.Write(buf[:n])
		if err != nil {
			log.Fatalf("cannot write to standard output: %s", err)
		}
	}
}
