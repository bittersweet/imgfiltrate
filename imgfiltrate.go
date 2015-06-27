package main

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
var dict []string

func loadDictionary(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Could not read dict")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()

		// Only save words that are longer than N character
		if len(word) > 2 {
			dict = append(dict, word)
		}
	}
}

// TODO: optimize algo
// Strips non alphabetical characters and compares them with our dict
// This means can't -> cant, it's -> its etc
// I've added those to a custom dict.txt untill I can find a better way
func isInDictionary(word string) bool {
	re := regexp.MustCompile("[^a-zA-Z]")
	word = re.ReplaceAllString(word, "")
	word = strings.ToLower(word)

	for _, dWord := range dict {
		if word == dWord {
			return true
		}
	}

	return false
}

// regular text has more alpha than non-alpha
func hasMoreAlphaCharacters(in []string) bool {
	alpha, nonAlpha := analyzeWords(in)
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

func processImage(path string) string {
	text := gosseract.Must(gosseract.Params{Src: path})
	text = strings.TrimSpace(text)
	words := strings.Fields(text)

	util.Log(fmt.Sprintf("Text: %s\n", text))

	flagged := false

	hasMoreAlpha := hasMoreAlphaCharacters(words)

	if hasMoreAlpha {
		util.Log("FLAGGING: hasMoreAlphaCharacters")
		flagged = true
	}

	for _, word := range words {
		if hasMoreAlpha {
			if isInDictionary(word) {
				util.Log(fmt.Sprintf("FLAGGING: isInDictionary: %s\n", word))
				flagged = true
			}
		}
	}

	if flagged {
		fmt.Printf("%s is BAD\n", path)
	} else {
		fmt.Printf("%s is GOOD\n", path)
	}

	return text
}

func main() {
	loadDictionary("/usr/share/dict/words")
	loadDictionary("/Users/markmulder/sources/go/src/github.com/bittersweet/ocrimages/dict.txt")

	processImage(os.Args[1])

	// images := util.LoadImagesFromDir("without")
	// for _, image := range images {
	// 	processImage(image)
	// }
}
