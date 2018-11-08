package main

import (
	"github.com/adelciotto/goprotein/goproteinpack"
)

func main() {
	packer, err := goproteinpack.NewPacker("./dna/test.txt")
	if err != nil {
		panic(err)
	}

	err = packer.SaveToFile("./dna/test.dna")
	if err != nil {
		panic(err)
	}
}
