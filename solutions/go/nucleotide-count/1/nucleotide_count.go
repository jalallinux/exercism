package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
// type Histogram ...
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
// type DNA ...
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (dna DNA) Counts() (Histogram, error) {
	// Initialize histogram with all valid nucleotides set to 0
	histogram := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	// Count each nucleotide in the DNA string
	for _, nucleotide := range dna {
		// Check if the nucleotide is valid (A, C, G, or T)
		if nucleotide != 'A' && nucleotide != 'C' && nucleotide != 'G' && nucleotide != 'T' {
			return nil, errors.New("invalid nucleotide")
		}
		histogram[nucleotide]++
	}

	return histogram, nil
}
