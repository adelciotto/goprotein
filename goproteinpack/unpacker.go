package goproteinpack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Unpacker struct {
	filename string
	file     *os.File
	stream   *Stream
}

const (
	numCodonsPerByte            = 4
	maxUnpackerStreamReadLength = 3
	numCodonsPerRead            = numCodonsPerByte * maxUnpackerStreamReadLength
)

var byteCodonMap = map[byte]byte{
	0: 'T',
	1: 'C',
	2: 'A',
	3: 'G',
}

func NewUnpacker(path string) (*Unpacker, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stream := NewStream(file, maxUnpackerStreamReadLength)

	return &Unpacker{filepath.Base(path), file, stream}, nil
}

// TODO: Remove this and replace with interface to the stream
// Another file will consume the unpacked stream and perform the DNA -> mRNA -> Protein translation
func (unpacker *Unpacker) Unpack(path string) error {
	outfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outfile.Close()
	defer unpacker.file.Close()

	writer := bufio.NewWriter(outfile)
	start := time.Now()
	for {
		bytes, err := unpacker.stream.Read()
		if err == io.EOF {
			break
		}

		for _, byteItem := range bytes {
			codons := unpackCodons(byteItem)
			writer.WriteString(string(codons[:]))
		}
	}
	writer.Flush()
	elapsed := time.Since(start)

	fmt.Printf("%s has been unpacked successfully to %s\n", unpacker.filename, path)
	fmt.Printf("unpacking took %s\n", elapsed)

	return nil
}

func unpackCodons(packed byte) []byte {
	codons := make([]byte, numCodonsPerByte)

	codons[0] = byteCodonMap[(packed>>6)&3]
	codons[1] = byteCodonMap[(packed>>4)&3]
	codons[2] = byteCodonMap[(packed>>2)&3]
	codons[3] = byteCodonMap[packed&3]

	return codons
}
