package dna

import (
	"errors"
)

type Nucleotide rune

const (
	A Nucleotide = 'A'
	C Nucleotide = 'C'
	G Nucleotide = 'G'
	T Nucleotide = 'T'
)

type Histogram map[Nucleotide]int

type DNA string

func isValidNucleotid(E rune) (bool, Nucleotide) {
	switch E {
	case 'A', 'C', 'G', 'T':
		return true, Nucleotide(E)
	default:
		return false, ' '
	}
}

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		A: 0,
		C: 0,
		G: 0,
		T: 0,
	}

	for _, v := range d {
		ok, nuc := isValidNucleotid(v)
		if !ok {
			return nil, errors.New("invalid input")
		}

		h[nuc]++
	}

	return h, nil
}
