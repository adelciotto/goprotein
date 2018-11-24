package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/adelciotto/goprotein/pkg/pack"
)

func main() {
	inputFilePtr := flag.String("input-file", "", "the input dna file")

	flag.Parse()

	if *inputFilePtr == "" {
		printError("input dna file not provided")
		printUsage()
		return
	}

	err := translateDNA(*inputFilePtr)
	if err != nil {
		printError(err.Error())
	}
}

// TODO: Actually translate DNA sequence into proteins
func translateDNA(inputfilepath string) error {
	inputFile, err := os.Open(inputfilepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	reader := pack.NewDNAReader(bufio.NewReader(inputFile))
	err = reader.ReadContents(func(codon string) {
		fmt.Print(codon)
	})
	if err != nil {
		return err
	}

	return nil
}

func printError(errMsg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s\n", errMsg))
}

func printUsage() {
	fmt.Println("usage: translate --input-file=<input-file>")
}
