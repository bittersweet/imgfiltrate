package color

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	_ "golang.org/x/image/webp"
	"io"
	"log"
)

func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func convertColorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return RGBToHex(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

func ProcessImage(contents io.Reader) (float64, int) {
	m, _, err := image.Decode(contents)
	if err != nil {
		log.Fatal("ProcessImage", err)
	}

	bounds := m.Bounds()
	total := 0
	colormap := make(map[color.Color]int)

	// fmt.Printf("min.y: %d max.y: %d", bounds.Min.Y, bounds.Max.Y)

	// We only care about the center 80%, so slice of 10% of the beginning and
	// end

	pctY := bounds.Max.Y / 100
	startY := 10 * pctY
	endY := 90 * pctY

	pctX := bounds.Max.X / 100
	startX := 10 * pctX
	endX := 90 * pctX

	for y := startY; y < endY; y++ {
		if y%2 == 0 {
			y++
		}

		for x := startX; x < endX; x++ {
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

	uniqueColors := len(colormap)

	percentage := float64(highest) / float64(total) * 100
	return percentage, uniqueColors
}
