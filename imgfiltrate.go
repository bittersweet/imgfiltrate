package main

import (
	"os"

	"github.com/bittersweet/imgfiltrate/ocr"
)

func main() {
	ocr.ProcessImage(os.Args[1])
	// images := util.LoadImagesFromDir("text")
	// for _, image := range images {
	// 	processImage(image)
	// }
}
