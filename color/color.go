package color

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
)

func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func convertColorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return RGBToHex(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

func ProcessImage(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	total := 0
	colormap := make(map[color.Color]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		if y%2 == 0 {
			y++
		}

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if x%2 == 0 {
				x++
			}
			total++
			color := m.At(x, y)

			colormap[color]++
		}
	}

	// Get color that occurs in the image the most
	highest := 0
	for _, amount := range colormap {
		if amount > highest {
			highest = amount
		}
	}

	percentage := float64(highest) / float64(total) * 100
	fmt.Println(percentage)
}
