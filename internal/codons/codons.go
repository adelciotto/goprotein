package codons

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
}
