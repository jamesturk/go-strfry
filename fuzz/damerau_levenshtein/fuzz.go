package damerau_levenshtein

import "github.com/jamesturk/go-jellyfish"
import "math/rand"
import "time"

func Fuzz(data []byte) int {
	jellyfish.DamerauLevenshtein(string(data), string(data))
	if len(data) > 2 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(data))
		jellyfish.DamerauLevenshtein(string(data[:n]), string(data[n:]))
	}
	return 0
}