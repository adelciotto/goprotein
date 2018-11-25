package translate

import (
	"fmt"

	"github.com/adelciotto/goprotein/internal/codons"
	"github.com/adelciotto/goprotein/pkg/pack"
)

type DNATranslator struct {
	dnaReader *pack.DNAReader
}

func NewDNATranslator(dnaReader *pack.DNAReader) *DNATranslator {
	return &DNATranslator{dnaReader}
}

func (translator *DNATranslator) Translate() error {
	return translator.dnaReader.ReadContents(func(codon string) {
		rnaCodon := codons.DNACodonToRNA(codon)
		protein, found := codons.RNACodonTable[rnaCodon]
		if found {
			fmt.Printf("%s: %s ", codon, protein)
		} else {
			fmt.Print("UNKNOWN")
		}
	})
}
