package translate

import (
	"fmt"
	"io"

	"github.com/adelciotto/goprotein/internal/codons"
	"github.com/adelciotto/goprotein/pkg/pack"
)

type DNATranslator struct {
	dnaReader *pack.DNAReader
	writer    io.Writer
}

func NewDNATranslator(dnaReader *pack.DNAReader, writer io.Writer) *DNATranslator {
	return &DNATranslator{dnaReader, writer}
}

func (translator *DNATranslator) Translate() error {
	return translator.dnaReader.ReadContents(func(codon string) {
		rnaCodon := codons.DNACodonToRNA(codon)
		protein, found := codons.RNACodonTable[rnaCodon]
		if found {
			output := fmt.Sprintf("%s: %s ", codon, protein)
			translator.writer.Write([]byte(output))
		} else {
			translator.writer.Write([]byte("UNKNOWN"))
		}
	})
}
