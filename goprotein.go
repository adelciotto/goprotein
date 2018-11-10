package main

import (
	"fmt"

	"github.com/adelciotto/goprotein/goproteinpack"
)

func main() {
	testFile := "./dna/test.txt"

	// TODO: Replace with real CLI program
	packer, err := goproteinpack.NewPacker(testFile)
	if err != nil {
		panic(err)
	}

	err = packer.Pack("./dna/test.rna")
	if err != nil {
		panic(err)
	}

	unpacker, err := goproteinpack.NewUnpacker("./dna/test.rna")
	if err != nil {
		panic(err)
	}

	fmt.Printf("RNA transcribed sequence for %s\n", testFile)

	buffer := make([]byte, 3)
	for {
		codons, err := unpacker.Read(buffer)
		if err != nil {
			break
		}

		for _, codon := range codons {
			fmt.Printf("%6s", codon)

			if codon == "UGA" {
				break
			}
		}
	}

	fmt.Println()
}
