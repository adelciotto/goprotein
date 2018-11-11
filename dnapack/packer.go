package dnapack

import (
	"fmt"
	"io"
	"unicode"
)

type Packer struct {
	stream *Stream
	writer io.Writer
}

const (
	readLength    = 4
	maxShiftWidth = 6
)

var nucleotideByteMap = map[byte]byte{
	'T': 0,
	'C': 1,
	'A': 2,
	'G': 3,
}

func NewPacker(reader io.Reader, writer io.Writer) *Packer {
	stream := NewStream(reader, readLength)
	return &Packer{stream, writer}
}

func (packer *Packer) Pack() error {
	err := packer.stream.ReadContents(func(nucleotides []byte) error {
		packedNucleotides, err := packer.packNucleotides(nucleotides)
		if err != nil {
			return err
		}

		packer.writer.Write([]byte{packedNucleotides})
		return nil
	})

	return err
}

func (packer *Packer) packNucleotides(nucleotides []byte) (byte, error) {
	var packedNucleotides byte

	for index, nucleotide := range nucleotides {
		nucleotide = byte(unicode.ToUpper(rune(nucleotide)))
		shiftWidth := byte(maxShiftWidth - (2 * index))
		nucleotideAsByte, ok := nucleotideByteMap[nucleotide]
		if !ok {
			return 0, fmt.Errorf("invalid nucleotide '%c' in DNA sequence", nucleotide)
		}

		if shiftWidth > 0 {
			packedNucleotides |= nucleotideAsByte << shiftWidth
		} else {
			packedNucleotides |= nucleotideAsByte
		}
	}

	return packedNucleotides, nil
}
