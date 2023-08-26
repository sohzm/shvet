package operations

import (
	"fmt"
	"image"


	"github.com/sz47/shvet/data"
	"github.com/sz47/shvet/structs"
)

type SColor struct {
	R int
	G int
	B int

	H int
	S int
	V float64  // brightness
}

func ReturnValue(r, g, b float64) float64 { 
	// Return Value from hsV
	var n float64
	if r > g && r > b {
		n = r
	}
	if g > r && g > b {
		n = g
	}
	n = b
	return float64((n*100)/255)
}


// Engine4 is my new attempt at writing something useful :(
//
// Steps taken from this video: https://youtu.be/epPOrHG6SnE
//
// # Planned steps:
//   - Get a color palatte of theme provided
//   - Create a gradient based on brightness of palatte
//   - Replace highlight-shadows gradient map of input image with the theme gradient
//   - Choose a desecntt opacity of top image
//   - TODO
//   - Blend new highlight-shadows-theme gradient map with the input image
//   - Get High - mid - low tones from input image
//   - Adjust the curves to match high-mid-low tones from image
//   - Adjust the vibrance/sturation/contrast whatever
func Engine4(structRGBA *structs.StructRGBA, themeName, blend string) {
	themeData := data.DataMap[themeName]
	themeGradient := make([]SColor, 0)
	for _, val := range themeData.SmallPalatte {
		temp := SColor {
			R: int(val[0]),
			G: int(val[1]),
			B: int(val[2]),
			V: ReturnValue(float64(val[0]), float64(val[1]), float64(val[2])),
		}
		themeGradient = append(themeGradient, temp)
	}
	fmt.Print(">> ", themeGradient)

	return
}



//	gradient := giveMeGradientForTheme(themeData.SmallPalatte)
//	brightnessMap := giveMeBrightnessMap(*structRGBA)
//	newBrightnessMap := giveMeBrightnessMapAndGradientMerged(gradient, brightnessMap)
//	opacity := findOpacityRequired(newBrightnessMap)
//	blended := blendEm(opacity, newBrightnessMap, structRGBA.Rgba)
//
//	changeEm(structRGBA.Rgba, blended)

func giveMeGradientForTheme(themeData [][]uint8) (colorx []structs.Colorx) {
	return
}

func giveMeBrightnessMap(rgba structs.StructRGBA) (colorx [][]structs.Colorx) {
	return
}

func giveMeBrightnessMapAndGradientMerged(gradient []structs.Colorx, brightnessMap [][]structs.Colorx) (colorx [][]structs.Colorx) {
	return
}

func findOpacityRequired(n [][]structs.Colorx) float32 {
	return 0.5
}

func blendEm(opacity float32, colorx [][]structs.Colorx, rgba *image.RGBA) (retColor [][]structs.Colorx) {
	return
}

func changeEm(rgba *image.RGBA, blended [][]structs.Colorx) {
	return
}
