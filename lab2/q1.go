package lab2

import (
	"os"
	"strings"
	"sync"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	s = strings.ToLower(s)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(s))
	for _, word := range s {
		go func(w string) {
			defer wg.Done()
			mu.Lock()
			words[w]++
			mu.Unlock()
		}(string(word))
	}
	return words
}

func WordCountPrint(fileName string) {
	sample, _ := os.ReadFile(fileName)
	sampleString := string(sample)

	m := WordCount(sampleString)
	for k, v := range m {
		if k == " " {
			k = "(blank)"
		}
		if k == "\n" {
			k = "(newline)"
		}
		println(k, ": ", v)
	}
}
