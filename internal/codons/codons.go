package codons

import (
	"strings"
)

const NucleotidesPerCodon = 3

var Nucleotides = []byte{'T', 'C', 'A', 'G'}

var DNAStopCodons = map[string]bool{
	"TGA": true,
	"TAA": true,
	"TAG": true,
}

var RNACodonTable = map[string]string{
	// Uracil
	"UUU": "F", "UCU": "S", "UAU": "Y", "UGU": "C",
	"UUC": "F", "UCC": "S", "UAC": "Y", "UGC": "C",
	"UUA": "L", "UCA": "S", "UAA": "STOP", "UGA": "STOP",
	"UUG": "L", "UCG": "S", "UAG": "STOP", "UGG": "W",

	// Cytosine
	"CUU": "L", "CCU": "P", "CAU": "H", "CGU": "R",
	"CUC": "L", "CCC": "P", "CAC": "H", "CGC": "R",
	"CUA": "L", "CCA": "P", "CAA": "Q", "CGA": "R",
	"CUG": "L", "CCG": "P", "CAG": "Q", "CGG": "R",

	// Adenine
	"AUU": "L", "ACU": "T", "AAU": "N", "AGU": "S",
	"AUC": "L", "ACC": "T", "AAC": "N", "AGC": "S",
	"AUA": "L", "ACA": "T", "AAA": "K", "AGA": "R",
	"AUG": "MET", "ACG": "T", "AAG": "K", "AGG": "R",

	// Guanine
	"GUU": "V", "GCU": "A", "GAU": "D", "GGU": "G",
	"GUC": "V", "GCC": "A", "GAC": "D", "GGC": "G",
	"GUA": "V", "GCA": "A", "GAA": "E", "GGA": "G",
	"GUG": "V", "GCG": "A", "GAG": "E", "GGG": "G",
}

func DNACodonToRNA(codon string) string {
	return strings.Replace(codon, "T", "U", NucleotidesPerCodon)
}
