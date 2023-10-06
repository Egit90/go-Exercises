package strand

func ToRNA(dna string) string {
	RNA := ""
	for _, v := range dna {
		switch v {
		case 'G':
			RNA += "C"
		case 'C':
			RNA += "G"
		case 'T':
			RNA += "A"
		case 'A':
			RNA += "U"
		}
	}
	return RNA
}
