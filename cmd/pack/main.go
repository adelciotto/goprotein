package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/adelciotto/goprotein/pkg/pack"
)

func main() {
	var outputFile string
	inputFilePtr := flag.String("input-file", "", "the input text file")
	flag.StringVar(&outputFile, "output-file", "", "the output dna file")

	flag.Parse()

	if *inputFilePtr == "" {
		printError("input file not provided")
		printUsage()
		return
	}

	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.dna", filenameWithoutExt(*inputFilePtr))
	}

	err := packDna(*inputFilePtr, outputFile)
	if err != nil {
		printError(err.Error())
	}
}

func packDna(inputFilepath string, outputFilepath string) error {
	inputFile, err := os.Open(inputFilepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilepath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	packer := pack.NewDNAPacker(bufio.NewReader(inputFile), writer)

	start := time.Now()
	err = packer.Pack()
	if err != nil {
		return err
	}
	writer.Flush()
	elapsed := time.Since(start)
	fmt.Printf("%s packed to %s in %s\n", inputFilepath, outputFilepath, elapsed)

	inputFileStat, err := inputFile.Stat()
	if err != nil {
		return err
	}
	outputFileStat, err := outputFile.Stat()
	if err != nil {
		return err
	}
	fmt.Printf("%s = %d bytes, %s = %d bytes\n", inputFilepath, inputFileStat.Size(), outputFilepath, outputFileStat.Size())

	return nil
}

func filenameWithoutExt(filepath string) string {
	return strings.TrimSuffix(filepath, path.Ext(filepath))
}

func printError(errMsg string) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s\n", errMsg))
}

func printUsage() {
	fmt.Println("usage: pack --input-file=<input-file> --output-file=<output-file>")
}
