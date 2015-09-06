package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func execute(prog []Instruction, in io.Reader, out io.Writer) error {
	var pc int
	var ptr int
	data := make([]byte, 30000)

	for {
		if pc >= len(prog) {
			return nil
		}

		i := prog[pc]
		branch := false

		switch i.Name() {
		case "INCD":
			ptr++
		case "DECD":
			ptr--
		case "INCP":
			data[ptr] = data[ptr] + 1
		case "DECP":
			data[ptr] = data[ptr] - 1
		case "OUT":
			b := data[ptr : ptr+1]
			out.Write(b)
		case "IN":
			b := data[ptr : ptr+1]
			in.Read(b)
		case "BZ":
			if data[ptr] == 0 {
				pc += i.Immediate()
				branch = true
			}
		case "BN":
			if data[ptr] != 0 {
				pc -= i.Immediate()
				branch = true
			}
		default:
			return fmt.Errorf("Unknown instruction: %s", i.Name())
		}

		if !branch {
			pc++
		}
	}
}

func TestHello(t *testing.T) {
	hello, err := ioutil.ReadFile("hello.bf")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	stripped := strip(hello)
	instructions := assemble(stripped)

	var in bytes.Buffer
	var out bytes.Buffer
	if err := execute(instructions, &in, &out); err != nil {
		t.Errorf("Failed to execute: %v", err)
	}

	want := "Hello World!\n"
	got := out.String()
	if got != want {
		t.Errorf("Output got %q want %q", got, want)
	}
}

func TestMandelbrot(t *testing.T) {
	mandelbrot, err := ioutil.ReadFile("mandelbrot.bf")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	stripped := strip(mandelbrot)
	instructions := assemble(stripped)

	var in bytes.Buffer
	var out bytes.Buffer
	if err := execute(instructions, &in, &out); err != nil {
		t.Errorf("Failed to execute: %v", err)
	}

	want, err := ioutil.ReadFile("mandelbrot.out")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	got := out.String()
	if got != string(want) {
		t.Errorf("Output got %q want %q", got, want)
	}
}
