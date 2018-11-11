package dnapack

import "io"

type Stream struct {
	reader io.Reader
	buffer []byte
}

func NewStream(reader io.Reader, readLength int) *Stream {
	return &Stream{reader, make([]byte, readLength)}
}

func (stream *Stream) ReadContents(fn func([]byte) error) error {
	for {
		numBytesRead, err := stream.reader.Read(stream.buffer)
		if err != nil {
			if err == io.EOF {
				return fn(stream.buffer[:numBytesRead])
			}

			return err
		}

		fnErr := fn(stream.buffer[:numBytesRead])
		if fnErr != nil {
			return fnErr
		}
	}
}
