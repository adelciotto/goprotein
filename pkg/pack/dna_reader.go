package pack

import (
	"bytes"
	"io"

	"github.com/adelciotto/goprotein/internal/codons"
)

type DNAReader struct {
	stream *Stream
}

const (
	streamReadLength    = 3
	nucleotidesPerCodon = 3
)

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
		if index > 0 && (index+1)%nucleotidesPerCodon == 0 {
			codon := buffer.String()
			result = append(result, codon)

			if _, match := codons.DNAStopCodons[codon]; match {
				break
			}

			buffer.Reset()
		}
	}

	return result
}

func unpackNucleotides(packedNucleotides []byte) string {
	var buffer bytes.Buffer

	for _, packedNucleotidesByte := range packedNucleotides {
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>6)&nucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>4)&nucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[(packedNucleotidesByte>>2)&nucleotidesPerCodon])
		buffer.WriteByte(byteToNucleotideMap[packedNucleotidesByte&nucleotidesPerCodon])
	}

	return buffer.String()
}
