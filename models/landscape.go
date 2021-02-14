package models

import (
	"../calculations/perlin2d"
	"math"
)

type Landscape struct {
	Width, Height int
	Map [][]int
}

func NewLandscape(height, width int) (*Landscape, error) {
	var landscape Landscape

	landscape.Height = height
	landscape.Width = width

	landscape.Map = make([][]int, width)
	for i := range landscape.Map {
		landscape.Map[i] = make([]int, height)
	}

	return &landscape, nil
}

func (landscape *Landscape) PerlinNoiseGeneration(scale float64, octaves int, persistence float64) *Landscape {
	perlin, _ := perlin2d.NewPerlin2D()

	for i := 0; i < landscape.Width; i++ {
		for j := 0; j < landscape.Height; j++ {
			vector := perlin2d.NewVector2D(i, j)
			landscape.Map[i][j] = (int)(math.Floor(scale * perlin.MultiOctaveNoise(vector, octaves, persistence) + scale / 2))
		}
	}

	return landscape
}