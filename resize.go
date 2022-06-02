package main

import (
	"image"
	"math"

	"golang.org/x/image/draw"
)

func Resize(img image.Image, ratio int) image.Image {
	newWidth := int(math.Ceil(float64(img.Bounds().Max.X) / 100 * float64(ratio)))
	newHeight := int(math.Ceil(float64(img.Bounds().Max.Y) / 100 * float64(ratio)))
	result := image.NewRGBA(
		image.Rect(
			0,
			0,
			newWidth,
			newHeight,
		),
	)
	// resize
	draw.BiLinear.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)

	return result
}
