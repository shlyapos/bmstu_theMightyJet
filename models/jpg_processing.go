package models

import (
	"github.com/fogleman/gg"
)

func (landscape *Landscape) getMaxHeight() int {
	maxHeight := 0

	for _, row := range landscape.Map {
		for _, elem := range row {
			if elem > maxHeight {
				maxHeight = elem
			}
		}
	}

	return maxHeight
}

func HeightToRGB(maxHeight, height int) float64 {
	tone := height * 256 / maxHeight
	return float64(tone)
}

func (landscape *Landscape) ToJPG(path string, quality int) {
	width, height := landscape.Width, landscape.Height

	context := gg.NewContext(width, height)
	context.SetRGB(0, 0, 0)
	context.Clear()

	max := landscape.getMaxHeight()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			currentColor := HeightToRGB(max, landscape.Map[x][y])
			context.SetRGB(currentColor, currentColor, currentColor)
			context.SetPixel(x, y)
		}
	}

	context.SavePNG(path)
}


