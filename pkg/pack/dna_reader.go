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
		buffer.WriteByte(codons.Nucleotides[(packedNucleotidesByte>>6)&codons.NucleotidesPerCodon])
		buffer.WriteByte(codons.Nucleotides[(packedNucleotidesByte>>4)&codons.NucleotidesPerCodon])
		buffer.WriteByte(codons.Nucleotides[(packedNucleotidesByte>>2)&codons.NucleotidesPerCodon])
		buffer.WriteByte(codons.Nucleotides[packedNucleotidesByte&codons.NucleotidesPerCodon])
	}

	return buffer.String()
}
