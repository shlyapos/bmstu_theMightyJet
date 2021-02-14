package jpg

import (
	_ "../../models"
	"image"
	"image/draw"
)

type ImageHandler struct {
	Width, Height int
	Image draw.Image
}

func NewImageHandler(width, height int) *ImageHandler {
	var handler ImageHandler

	handler.Width = width
	handler.Height = height
	img := image.Rect(0, 0, width, height)
	handler.Image =

	return &handler
}

