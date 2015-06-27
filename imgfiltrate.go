package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bittersweet/imgfiltrate/color"
	"github.com/bittersweet/imgfiltrate/ocr"
)

type Result struct {
	ColorPercentage         float64 `json:"color_percentage"`
	TotalColors             float64 `json:"total_colors"`
	AlphabeticCharacters    int     `json:"alphabetic_characters"`
	NonAlphabeticCharacters int     `json:"non_alphabetic_characters"`
	TotalCharacters         int     `json:"total_characters"`
	DictionaryWords         int     `json:"dictionary_words"`
	Advice                  string  `json:"advice"`
}

func processImage(file string) Result {
	alpha, nonAlpha, totalCharacters, dictionaryWords := ocr.ProcessImage(file)
	pct, totalColors := color.ProcessImage(file)

	r := Result{
		ColorPercentage:         pct,
		TotalColors:             totalColors,
		AlphabeticCharacters:    alpha,
		NonAlphabeticCharacters: nonAlpha,
		TotalCharacters:         totalCharacters,
		DictionaryWords:         dictionaryWords,
	}

	return r
}

func (r *Result) advise() {
	if r.DictionaryWords > 0 {
		r.Advice = "BAD"
		return
	}

	if r.AlphabeticCharacters > r.NonAlphabeticCharacters && r.AlphabeticCharacters > 10 {
		r.Advice = "BAD"
		return
	}

	// TODO: Take ColorPercentage into account

	r.Advice = "GOOD"
}

func (r *Result) output() {
	output, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal("MarshalIndent", err)
	}

	fmt.Println(string(output))
}

func main() {
	file := os.Args[1]
	r := processImage(file)
	r.advise()
	r.output()

	// images := util.LoadImagesFromDir("without")
	// for _, image := range images {
	// 	fmt.Println("processing", image)
	// 	r := processImage(image)
	// 	r.advise()
	// 	r.output()
	// }
}
