package ocr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	loadDictionary("/usr/share/dict/words")
}

func TestIsInDictionary(t *testing.T) {
	outcome := isInDictionary("test")
	assert.Equal(t, true, outcome)

	outcome = isInDictionary("bogusasdasd")
	assert.Equal(t, false, outcome)
}

// Run with go test -v -bench=.
func BenchmarkIsInDictionary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isInDictionary("test")
	}
}
