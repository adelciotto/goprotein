package goproteinpack

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
)

type Unpacker struct {
	filename   string
	file       *os.File
	fileReader *bufio.Reader
}

const (
	numCodonsPerRead            = numCodonsPerByte * maxUnpackerStreamReadLength
	numCodonsPerByte            = 4
	maxUnpackerStreamReadLength = 3
)

var byteCodonMap = map[byte]byte{
	0: 'U',
	1: 'C',
	2: 'A',
	3: 'G',
}

func NewUnpacker(path string) (*Unpacker, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	return &Unpacker{filepath.Base(path), file, reader}, nil
}

func (unpacker *Unpacker) Read(buffer []byte) ([]string, error) {
	_, err := unpacker.fileReader.Read(buffer)
	if err != nil {
		return nil, err
	}

	var stringBuffer bytes.Buffer

	for _, packed := range buffer {
		codons := unpackCodons(packed)
		stringBuffer.WriteString(codons)
	}
	codonsString := stringBuffer.String()
	stringBuffer.Reset()

	var result []string
	for index, code := range codonsString {
		stringBuffer.WriteRune(code)
		if index > 0 && (index+1)%codesPerCodon == 0 {
			result = append(result, stringBuffer.String())
			stringBuffer.Reset()
		}
	}

	return result, nil
}

func unpackCodons(packed byte) string {
	var codons bytes.Buffer

	codons.WriteByte(byteCodonMap[(packed>>6)&3])
	codons.WriteByte(byteCodonMap[(packed>>4)&3])
	codons.WriteByte(byteCodonMap[(packed>>2)&3])
	codons.WriteByte(byteCodonMap[packed&3])

	return codons.String()
}
