package ocr

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bittersweet/imgfiltrate/util"

	"github.com/otiai10/gosseract"
)

// Global var that holds the dictionary
var dict = make(map[string]string)

func loadDictionary(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Could not read dict")
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		word := scanner.Text()
		word = processWord(word)
		// Only store words that are longer than 2 characters
		if len(word) > 2 {
			dict[word] = ""
		}
	}
}

func isInDictionary(word string) bool {
	word = processWord(word)

	_, ok := dict[word]
	return ok
}

func processWord(word string) string {
	re := regexp.MustCompile("[^a-zA-Z]")
	word = re.ReplaceAllString(word, "")
	word = strings.ToLower(word)
	return word
}

// regular text has more alpha than non-alpha
func hasMoreAlphaCharacters(alpha int, nonAlpha int) bool {
	if alpha > nonAlpha {
		return true
	}

	return false
}

func analyzeWords(words []string) (int, int) {
	alpha := 0
	nonAlpha := 0

	for _, word := range words {
		re := regexp.MustCompile("[a-zA-Z]")
		for _, c := range word {
			comp := []byte(string(c))
			if re.Match(comp) {
				alpha++
			} else {
				nonAlpha++
			}
		}
	}
	util.Log(fmt.Sprintf("Alpha: %d NonAlpha: %d\n", alpha, nonAlpha))
	return alpha, nonAlpha
}

func ProcessImage(path string) (int, int, int, int) {
	loadDictionary("/usr/share/dict/words")

	text := gosseract.Must(gosseract.Params{Src: path})
	totalCharacters := len(text)
	text = strings.TrimSpace(text)
	words := strings.Fields(text)

	util.Log(fmt.Sprintf("Text: %s\n", text))

	alpha, nonAlpha := analyzeWords(words)
	hasMoreAlpha := hasMoreAlphaCharacters(alpha, nonAlpha)

	if hasMoreAlpha {
		util.Log("hasMoreAlphaCharacters")
	}

	dictionaryWords := 0
	for _, word := range words {
		if hasMoreAlpha {
			if isInDictionary(word) {
				dictionaryWords++
				util.Log(fmt.Sprintf("isInDictionary: %s\n", word))
			}
		}
	}

	return alpha, nonAlpha, totalCharacters, dictionaryWords
}
