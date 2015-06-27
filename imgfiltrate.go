package main

import (
	"os"

	"github.com/bittersweet/imgfiltrate/color"
	"github.com/bittersweet/imgfiltrate/ocr"
)

func main() {
	file := os.Args[1]
	ocr.ProcessImage(file)
	color.ProcessImage(file)
	// images := util.LoadImagesFromDir("text")
	// for _, image := range images {
	// 	processImage(image)
	// }
}
