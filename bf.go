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

// strip strips the program of all invalid commands.
func strip(prog []byte) (out []BfInst) {
	for _, inst := range prog {
		switch BfInst(inst) {
		case IncD:
			fallthrough
		case DecD:
			fallthrough
		case IncP:
			fallthrough
		case DecP:
			fallthrough
		case Out:
			fallthrough
		case In:
			fallthrough
		case JmpF:
			fallthrough
		case JmpB:
			out = append(out, BfInst(inst))
		}
	}
	return
}

func emit(prog []BfInst) {
	for _, inst := range prog {
		switch inst {
		case IncD:
			log.Println("INCD")
		case DecD:
			log.Println("DECD")
		case IncP:
			log.Println("INCP")
		case DecP:
			log.Println("DECP")
		case Out:
			log.Println("OUT")
		case In:
			log.Println("IN")
		case JmpF:
			log.Println("BZ TODO")
		case JmpB:
			log.Println("BN TODO")
		default:
			log.Fatalf("Bad instrustion: '%c'", inst)
		}
	}
}

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

	input, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("Unable to read program: %v", err)
		os.Exit(1)
	}

	log.Printf("Program: %s\n", input)

	stripped := strip(input)
	log.Printf("Stripped program: %s\n", stripped)

	emit(stripped)
}
