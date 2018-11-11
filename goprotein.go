package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adelciotto/goprotein/dnapack"
)

const (
	dnaTxtFile    = "./data/test.txt"
	dnaPackedFile = "./data/test.dna"
)

func main() {
	// TODO: Replace with real CLI program
	infile, err := os.Open(dnaTxtFile)
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	outfile, err := os.Create(dnaPackedFile)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)
	packer := dnapack.NewPacker(bufio.NewReader(infile), writer)

	start := time.Now()
	err = packer.Pack()
	if err != nil {
		panic(err)
	}
	writer.Flush()
	elapsed := time.Since(start)
	fmt.Printf("%s packed to %s in %s\n", dnaTxtFile, dnaPackedFile, elapsed)

	infile, _ = os.Open(dnaPackedFile)

	unpacker := dnapack.NewUnpacker(bufio.NewReader(infile), os.Stdout)
	start = time.Now()
	err = unpacker.Unpack()
	if err != nil {
		panic(err)
	}
	elapsed = time.Since(start)
	fmt.Printf("\n%s unpacked to STDOUT in %s\n", dnaPackedFile, elapsed)
}
