package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("ErrStop")
	ErrInvalidBase = errors.New("ErrInvalidBase")
)

var pr map[string]string = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	myCondons := RnaToCodons(rna)
	myRes := []string{}
	for _, v := range myCondons {

		value, err := FromCodon(v)

		if err != nil && err != ErrStop {
			return nil, err
		}

		if err == ErrStop {
			return myRes, nil
		}

		myRes = append(myRes, value)

	}
	return myRes, nil
}

func FromCodon(codon string) (string, error) {
	v, ok := pr[codon]

	if !ok {
		return "", ErrInvalidBase
	}
	if ok && v == "STOP" {
		return "", ErrStop
	}

	return v, nil

}

func RnaToCodons(rna string) []string {
	res := []string{}
	for i := 0; i < len(rna); i += 3 {
		end := i + 3

		if end > len(rna) {
			end = len(rna)
		}

		res = append(res, rna[i:end])
	}

	return res
}
