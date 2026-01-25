package strand

var dnaToRNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna string
	for _, v := range dna {
		rna += string(dnaToRNA[v])
	}
	return rna
}
