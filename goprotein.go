package main

import (
	"github.com/adelciotto/goprotein/goproteinpack"
)

func main() {
	// TODO: Replace with real CLI program
	packer, err := goproteinpack.NewPacker("./dna/test.txt")
	if err != nil {
		panic(err)
	}

	err = packer.Pack("./dna/test.dna")
	if err != nil {
		panic(err)
	}

	unpacker, err := goproteinpack.NewUnpacker("./dna/test.dna")
	if err != nil {
		panic(err)
	}

	err = unpacker.Unpack("./dna/test-unpacked.txt")
	if err != nil {
		panic(err)
	}
}
