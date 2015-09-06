package main

import (
	"fmt"
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

type Instruction interface {
	Emit() string

	Name() string

	Immediate() int
}

type GeneralInst string

func (i GeneralInst) Emit() string {
	return string(i)
}

func (i GeneralInst) Name() string {
	return string(i)
}

func (i GeneralInst) Immediate() int {
	return 0
}

type Branch struct {
	// addr is the address of this branch
	addr int
	// name is BZ or BN.
	name string
	// dest is the branch destination instruction (pc-relative)
	dest int
}

func (b *Branch) Emit() string {
	return fmt.Sprintf("%s %d", b.name, b.dest)
}

func (b *Branch) Name() string {
	return b.name
}

func (b *Branch) Immediate() int {
	return b.dest
}

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

func assemble(prog []BfInst) (out []Instruction) {
	var pc int
	var branchStack []*Branch

	for _, inst := range prog {
		switch inst {
		case IncD:
			out = append(out, GeneralInst("INCD"))
		case DecD:
			out = append(out, GeneralInst("DECD"))
		case IncP:
			out = append(out, GeneralInst("INCP"))
		case DecP:
			out = append(out, GeneralInst("DECP"))
		case Out:
			out = append(out, GeneralInst("OUT"))
		case In:
			out = append(out, GeneralInst("IN"))
		case JmpF:
			b := Branch{
				addr: pc,
				name: "BZ",
			}
			branchStack = append(branchStack, &b)
			out = append(out, &b)
		case JmpB:
			dest := branchStack[len(branchStack)-1]
			branchStack = branchStack[:len(branchStack)-1]
			b := Branch{
				addr: pc,
				name: "BN",
				dest: pc - (dest.addr + 1),
			}
			dest.dest = (pc - dest.addr) + 1
			out = append(out, &b)
		default:
			log.Fatalf("Bad instrustion: '%c'", inst)
		}

		pc++
	}

	return
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

	instructions := assemble(stripped)

	for pc, i := range instructions {
		log.Printf("%d: %v", pc, i.Emit())
	}
}
