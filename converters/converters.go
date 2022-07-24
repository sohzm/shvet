package converters

import (
	"image"
	"image/draw"
	"shvet/structs"
)

func ImageToRGBA(img image.Image) (structs.StructRGBA, error) {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y

	rgba, ok := img.(*image.RGBA)
	if !ok {
		b := img.Bounds()
		rgba = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	}

	return structs.StructRGBA{
		Rgba: rgba,
		MaxX: maxX,
		MaxY: maxY,
	}, nil
}
