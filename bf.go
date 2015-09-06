package main

import (
	"io"
	"log"
	"os"
)

type BfInst byte

const (
	IncD BfInst = '>'
	DecD        = '<'
	IncP        = '+'
	DecP        = '-'
	Out         = '.'
	In          = ','
	JmpF        = '['
	JmpB        = ']'
)

func main() {
	var prog io.Reader

	if len(os.Args) < 2 {
		prog = os.Stdin
	} else {
		var err error
		prog, err = os.Open(os.Args[1])
		if err != nil {
			log.Printf("Error opening file: %v", err)
			os.Exit(1)
		}
	}

	b := make([]byte, 512)
	prog.Read(b)
	log.Printf("%s", b)
}
