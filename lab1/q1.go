package lab1

import (
	"math/rand"
)

var dnaLetters = []rune("CAGT")

func GenerateDNASample(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = dnaLetters[rand.Intn(len(dnaLetters))]
	}
	return string(b)
}

func HammingDistance(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	dist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}

/*
func main() {
	for i := 0; i < 1000; i++ {
		length := 20
		dna1 := GenerateDNASample(length)
		dna2 := GenerateDNASample(length)
		dist := HammingDistance(dna1, dna2)
		if dist != -1 {
			fmt.Println("Pair: ", (i + 1))
			fmt.Println("DNA 1: ", dna1)

			fmt.Println("DNA 2: ", dna2)
			fmt.Println("Harming distance: ", dist)

		} else {
			fmt.Println("DNA sequences are not of equal length")
		}
	}
}
*/
