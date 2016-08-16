package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bittersweet/imgfiltrate/color"
)

type Result struct {
	ColorPercentage float64 `json:"color_percentage"`
	TotalColors     int     `json:"total_colors"`
}

func processImage(file string) Result {
	pct, totalColors := color.ProcessImage(file)

	r := Result{
		ColorPercentage: pct,
		TotalColors:     totalColors,
	}

	return r
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
	r.output()
}
