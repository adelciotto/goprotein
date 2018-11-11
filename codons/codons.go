package codons

var stopCodons = map[string]bool{
	"TGA": true,
	"TAA": true,
	"TAG": true,
}

func IsStopCodon(codon string) bool {
	_, ok := stopCodons[codon]
	return ok
}
