package _kyu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func DNAStrand(dna string) (res string) {
	for _, s := range dna {
		switch s {
		case 65:
			s = 84
		case 67:
			s = 71
		case 71:
			s = 67
		case 84:
			s = 65
		}
		res += string(s)
	}
	return
}


func TestComplementaryDNA(t *testing.T) {
	assert.Equal(t, DNAStrand("AAAA"), "TTTT")
	assert.Equal(t, DNAStrand("ATTGC"), "TAACG")
	assert.Equal(t, DNAStrand("GTAT"), "CATA")
}
