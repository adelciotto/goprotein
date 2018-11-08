package goproteinpack

import (
	"bufio"
	"io"
	"os"
)

type Stream struct {
	reader        *bufio.Reader
	maxReadLength int
}

func NewStream(file *os.File, maxReadLength int) *Stream {
	reader := bufio.NewReader(file)
	return &Stream{reader, maxReadLength}
}

func (stream *Stream) Read() ([]byte, error) {
	bytes, _ := stream.reader.Peek(stream.maxReadLength)

	// remove any \n characters from the buffer
	for index, byteItem := range bytes {
		if byteItem == '\n' {
			bytes = append(bytes[:index], bytes[index+1:]...)
			break
		}
	}

	stream.reader.Discard(stream.maxReadLength)

	var err error
	if len(bytes) == 0 {
		err = io.EOF
	}
	return bytes, err
}
