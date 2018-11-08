package goproteinpack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"unicode"
)

type Packer struct {
	filename string
	file     *os.File
	stream   *Stream
}

const (
	maxPackerStreamReadLength = 4
	maxShiftWidth             = 6
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

	stream := NewStream(file, maxPackerStreamReadLength)

	return &Packer{filepath.Base(path), file, stream}, nil
}

func (packer *Packer) Pack(path string) error {
	outfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outfile.Close()
	defer packer.file.Close()

	writer := bufio.NewWriter(outfile)
	start := time.Now()
	for {
		bytes, err := packer.stream.Read()
		if err == io.EOF {
			break
		}

		codons := packCodons(bytes)
		writer.WriteByte(codons)
	}
	writer.Flush()
	elapsed := time.Since(start)

	fmt.Printf("%s has been packed successfully to %s\n", packer.filename, path)
	fmt.Printf("packing took %s\n", elapsed)

	return nil
}

func packCodons(codons []byte) byte {
	var result byte

	for index, data := range codons {
		codon := byte(unicode.ToUpper(rune(data)))
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
