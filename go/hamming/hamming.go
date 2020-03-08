package hamming

import (
	"fmt"
)

type DNAStrand string

type DNA struct {
	First  DNAStrand
	Second DNAStrand
}

func newDNA(firstDNA DNAStrand, secondDNA DNAStrand) DNA {
	return DNA{
		First:  firstDNA,
		Second: secondDNA,
	}
}

func Distance(firstDNA, secondDNA DNAStrand) (int, error) {
	dna := newDNA(firstDNA, secondDNA)

	if !dna.valid(){
		return 0, fmt.Errorf("DNA validation failed, Length did not match")
	}

	return dna.findDistance()
}

func (dna DNA)findDistance() (int, error) {
	distance := 0
	for i := range dna.Second {
		if dna.Second[i] != dna.First[i] {
			distance += 1
		}
	}
	return distance, nil
}

func (dna DNA) valid() bool {
	return len(dna.First) == len(dna.Second)
}
