package main

import (
	"fmt"
	"os"

	"github.com/bittersweet/imgfiltrate/color"
	"github.com/bittersweet/imgfiltrate/ocr"
)

type Result struct {
	ColorPercentage         float64
	TotalColors             float64
	AlphabeticCharacters    int
	NonAlphabeticCharacters int
	TotalCharacters         int
	DictionaryWords         int
	Advice                  string
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
	fmt.Printf("%#v\n", r)
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
