package main

import (
	"image"
	"image/draw"
)

func Merge(img1 image.Image, img2 image.Image) image.Image {
	img2Size := img2.Bounds().Size()
	// create starting point for img2
	startPoint2 := image.Point{img1.Bounds().Dx(), 0}
	// create img2 rect
	rect2 := image.Rect(
		startPoint2.X,
		startPoint2.Y,
		startPoint2.X+img2Size.X,
		startPoint2.Y+img2Size.Y,
	)
	// create rect for img1+img2
	rect := image.Rect(
		0,
		0,
		rect2.Max.X,
		rect2.Max.Y,
	)

	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, rect2, img2, image.Point{0, 0}, draw.Src)

	return rgba
}
