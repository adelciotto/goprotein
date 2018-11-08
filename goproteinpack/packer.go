package goproteinpack

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Packer struct {
	filename   string
	file       *os.File
	fileReader *bufio.Reader
}

const (
	codonsPerByte = 4
	maxShiftWidth = 6
)

var codonByteMap = map[byte]byte{
	'T': 0,
	'C': 1,
	'A': 2,
	'G': 3,
}

func NewPacker(path string) (*Packer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	return &Packer{filepath.Base(path), file, reader}, nil
}

func (p *Packer) SaveToFile(path string) error {
	outfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outfile.Close()
	defer p.file.Close()

	writer := bufio.NewWriter(outfile)

	for codons := p.readNextCodons(); len(codons) != 0; codons = p.readNextCodons() {
		packedCodons := packCodons(codons)
		fmt.Printf("writing: %08b to file %s\n", packedCodons, path)
		writer.WriteByte(packedCodons)
	}

	writer.Flush()
	return nil
}

func (p *Packer) readNextCodons() []byte {
	bytes, _ := p.fileReader.Peek(codonsPerByte)
	p.fileReader.Discard(codonsPerByte)

	return bytes
}

func packCodons(codons []byte) byte {
	var result byte

	for index, codon := range codons {
		shiftWidth := byte(maxShiftWidth - (2 * index))
		codonByte := codonByteMap[codon]

		if shiftWidth > 0 {
			result |= codonByte << shiftWidth
		} else {
			result |= codonByte
		}
	}

	return result
}
