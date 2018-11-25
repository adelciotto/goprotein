package pack

import (
	"bytes"
	"io"

	"github.com/adelciotto/goprotein/internal/codons"
)

type DNAReader struct {
	stream *Stream
}

const streamReadLength = 3

var byteToNucleotideMap = map[byte]byte{
	0: 'T',
	1: 'C',
	2: 'A',
	3: 'G',
}

func NewDNAReader(reader io.Reader) *DNAReader {
	stream := NewStream(reader, streamReadLength)
	return &DNAReader{stream}
}

func (reader *DNAReader) ReadContents(fn func(codon string)) error {
	return reader.stream.ReadContents(func(packedNucleotides []byte) error {
		unpackedNucleotides := unpackNucleotides(packedNucleotides)
		codons := splitNucleotidesIntoCodons(unpackedNucleotides)

		for _, codon := range codons {
			fn(codon)
		}

		return nil
	})
}

func splitNucleotidesIntoCodons(nucleotides string) []string {
	var result []string
	var buffer bytes.Buffer

	for index, nucleotide := range nucleotides {
		buffer.WriteRune(nucleotide)
		if index > 0 && (index+1)%codons.NucleotidesPerCodon == 0 {
			result = append(result, buffer.String())
			buffer.Reset()
		}
	}

	return result
}

func unpackNucleotides(packedNucleotides []byte) string {
	var buffer bytes.Buffer

	for _, packedNucleotidesByte := range packedNucleotides {
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>6)&codons.NucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>4)&codons.NucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>2)&codons.NucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[packedNucleotidesByte&codons.NucleotidesPerCodon])
	}

	return buffer.String()
}
