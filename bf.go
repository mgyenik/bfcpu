package main

import (
	"io"
	"io/ioutil"
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
	var r io.Reader

	if len(os.Args) < 2 {
		r = os.Stdin
	} else {
		var err error
		r, err = os.Open(os.Args[1])
		if err != nil {
			log.Printf("Error opening file: %v", err)
			os.Exit(1)
		}
	}

	prog, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("Unable to read program: %v", err)
		os.Exit(1)
	}

	log.Printf("%s", prog)
}
