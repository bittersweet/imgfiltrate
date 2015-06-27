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
	Advice                  string
}

func processImage(file string) Result {
	alpha, nonAlpha, total := ocr.ProcessImage(file)
	pct, totalColors := color.ProcessImage(file)

	r := Result{
		ColorPercentage:         pct,
		TotalColors:             totalColors,
		AlphabeticCharacters:    alpha,
		NonAlphabeticCharacters: nonAlpha,
		TotalCharacters:         total,
	}

	return r
}

func (r *Result) advise() {
	r.Advice = "BAD"
}

func (r *Result) output() {
	fmt.Printf("%#v\n", r)
}

func main() {
	file := os.Args[1]
	r := processImage(file)
	r.advise()
	r.output()

	// images := util.LoadImagesFromDir("text")
	// for _, image := range images {
	// 	processImage(image)
	// }
}
