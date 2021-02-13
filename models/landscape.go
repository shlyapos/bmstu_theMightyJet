package models

type Landscape struct {
	Width, Height int
}

func NewLandscape(height, width int) (*Landscape, error) {
	var landscape Landscape

	landscape.Height = height
	landscape.Width = width
	return &landscape, nil
}