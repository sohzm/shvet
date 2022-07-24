package operations

import (
	"image/color"
	"shvet/converters"
	"shvet/files"
)

func Darken(fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	for i := 0; i < structRGBA.MaxX; i++ {
		for j := 0; j < structRGBA.MaxY; j++ {
			r, g, b, _ := structRGBA.Rgba.At(i, j).RGBA()
			r_8 := uint8(r) / 2
			g_8 := uint8(g) / 2
			b_8 := uint8(b) / 2
			structRGBA.Rgba.Set(i, j, color.RGBA{r_8, g_8, b_8, 255})
		}
	}

	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

func Lighten(fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	for i := 0; i < structRGBA.MaxX; i++ {
		for j := 0; j < structRGBA.MaxY; j++ {
			r, g, b, _ := structRGBA.Rgba.At(i, j).RGBA()
			r_8 := uint8(r) * 2
			if r_8 < uint8(r) {
				r_8 = 255
			}
			g_8 := uint8(g) * 2
			if g_8 < uint8(g) {
				g_8 = 255
			}
			b_8 := uint8(b) * 2
			if b_8 < uint8(b) {
				b_8 = 255
			}
			structRGBA.Rgba.Set(i, j, color.RGBA{r_8, g_8, b_8, 255})
		}
	}

	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}
