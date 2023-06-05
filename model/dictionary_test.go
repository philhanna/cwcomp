package model

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDictionary_GetMatchingWords(t *testing.T) {
	pattern := `[^AEIOU][AEIOU].E`
	stime := time.Now()
	stop := make(chan struct{})
	have := make([]string, 0)
	type Stopper struct{}
	ch := GetMatchingWords(pattern, stop)
	for word := range ch {
		have = append(have, word)
		if len(have) > 10 {
			stop <- Stopper{}
			break
		}
	}
	close(stop)
	elapsed := time.Since(stime)
	fmt.Printf("%d words matched the pattern, elapsed time=%v\n", len(have), elapsed)
	if false { // Change to true if output is desired
		filename := filepath.Join(os.TempDir(), "have.txt")
		fp, _ := os.Create(filename)
		defer fp.Close()
		for _, word := range have {
			fmt.Fprintf(fp, "%s\n", word)
		}
	}
}
