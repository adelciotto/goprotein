package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/adelciotto/goprotein/pkg/pack"
	"github.com/adelciotto/goprotein/pkg/translate"
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

func translateDNA(inputfilepath string) error {
	inputFile, err := os.Open(inputfilepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	reader := pack.NewDNAReader(bufio.NewReader(inputFile))
	translator := translate.NewDNATranslator(reader, os.Stdout)
	err = translator.Translate()
	if err != nil {
		return err
	}

	fmt.Println()
	return nil
}

func printError(errMsg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s\n", errMsg))
}

func printUsage() {
	fmt.Println("usage: translate --input-file=<input-file>")
}
