//Package strand converts from DNA to RNA
package strand

//ToRNA takes in a DNA string and returns the RNA string
func ToRNA(dna string) string {
	rna := ""

	for i := 0; i < len(dna); i++ {
		if dna[i] == 'G' {
			rna += "C"
		} else if dna[i] == 'C' {
			rna += "G"
		} else if dna[i] == 'T' {
			rna += "A"
		} else if dna[i] == 'A' {
			rna += "U"
		}
	}

	return rna
}
