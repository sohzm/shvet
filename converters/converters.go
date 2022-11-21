package converters

import (
	"image"
	"image/draw"
	"github.com/sz47/shvet/structs"
)

func ImageToRGBA(img image.Image) (structs.StructRGBA, error) {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y

	rgba, ok := img.(*image.RGBA)
	if !ok {
		rgba = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
	}

	return structs.StructRGBA{
		Rgba: rgba,
		MaxX: maxX,
		MaxY: maxY,
	}, nil
}
