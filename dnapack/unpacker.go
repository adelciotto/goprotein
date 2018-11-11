package dnapack

import (
	"bytes"
	"io"

	"github.com/adelciotto/goprotein/codons"
)

type Unpacker struct {
	stream *Stream
	writer io.Writer
}

const (
	unpackerReadLength  = 3
	nucleotidesPerCodon = 3
)

var byteToNucleotideMap = map[byte]byte{
	0: 'T',
	1: 'C',
	2: 'A',
	3: 'G',
}

func NewUnpacker(reader io.Reader, writer io.Writer) *Unpacker {
	stream := NewStream(reader, unpackerReadLength)
	return &Unpacker{stream, writer}
}

func (unpacker *Unpacker) Unpack() error {
	return unpacker.UnpackWithFunc(func(codon string) {
		unpacker.writer.Write([]byte(codon))
	})
}

func (unpacker *Unpacker) UnpackWithFunc(fn func(codon string)) error {
	err := unpacker.stream.ReadContents(func(packedNucleotides []byte) error {
		nucleotides := unpacker.unpackNucleotides(packedNucleotides)
		codons := unpacker.splitNucleotidesIntoCodons(nucleotides)
		for _, codon := range codons {
			fn(codon)
		}

		return nil
	})

	return err
}

func (unpacker *Unpacker) splitNucleotidesIntoCodons(nucleotides string) []string {
	var result []string
	var buffer bytes.Buffer

	for index, nucleotide := range nucleotides {
		buffer.WriteRune(nucleotide)
		if index > 0 && (index+1)%nucleotidesPerCodon == 0 {
			codon := buffer.String()
			result = append(result, codon)

			if codons.IsStopCodon(codon) {
				break
			}

			buffer.Reset()
		}
	}

	return result
}

func (unpacker *Unpacker) unpackNucleotides(packedNucleotides []byte) string {
	var buffer bytes.Buffer

	for _, packedNucleotidesByte := range packedNucleotides {
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>6)&3])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>4)&3])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>2)&3])
		buffer.WriteByte(byteToNucleotideMap[packedNucleotidesByte&3])
	}

	return buffer.String()
}
