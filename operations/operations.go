package operations

import (
	"fmt"
	"github.com/sz47/shvet/data"
	"github.com/sz47/shvet/structs"
	"image/color"
	"sync"
)

// TODO make all engine functions take same data.Data struct
// itll make it eazy using it
// something like:
// Function (structRGBA *structs.structRGBA, data data.Data, args structs.Args)

func Engine(structRGBA *structs.StructRGBA, localTheme [][]uint8) {
	// rgb cube approach
	var extremes = [8][2]int{}
	// is there a better way to do this?
	for i := 0; i < 8; i++ {
		extremes[i] = [2]int{0, 0}
	}

	var wg sync.WaitGroup

	for i := 0; i < structRGBA.MaxX; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < structRGBA.MaxY; j++ {
				r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
				for k, val := range data.ImpPoints {
					tr, tg, tb, _ := toUint8(structRGBA.Rgba.At(extremes[k][0], extremes[k][1]).RGBA())
					magnitude := mag(r, g, b, val[0], val[1], val[2])
					tempMagnitude := mag(tr, tg, tb, val[0], val[1], val[2])
					if magnitude < tempMagnitude {
						extremes[k][0] = i
						extremes[k][1] = j
					}
				}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	var exValues = [8][3]uint8{}
	for i, val := range extremes {
		r, g, b, _ := toUint8(structRGBA.Rgba.At(val[0], val[1]).RGBA())
		exValues[i] = [3]uint8{r, g, b}
		//exValues[i] = data.ImpPoints[i]
	}

	for i := 0; i < structRGBA.MaxX; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < structRGBA.MaxY; j++ {
				r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
				var relR, relG, relB float64
				var finR, finG, finB uint8
				//diffWhite := (exValues[6][0] - r)*(exValues[6][0] - r) + (exValues[6][1] - g)*(exValues[6][1] - g) + (exValues[6][2] - b)*(exValues[6][2] - b)
				//diffBlack := (exValues[7][0] - r)*(exValues[7][0] - r) + (exValues[7][1] - g)*(exValues[7][1] - g) + (exValues[7][2] - b)*(exValues[7][2] - b)
				relR = float64((r - exValues[7][0])) / float64((exValues[6][0] - exValues[7][0]))
				relG = float64((g - exValues[7][1])) / float64((exValues[6][1] - exValues[7][1]))
				relB = float64((b - exValues[7][2])) / float64((exValues[6][2] - exValues[7][2]))
				finR = controlOverflow(int64(uint8(((float64(localTheme[6][0] - localTheme[7][0])) * relR))), int64(localTheme[7][0]))
				finG = controlOverflow(int64(uint8(((float64(localTheme[6][1] - localTheme[7][1])) * relG))), int64(localTheme[7][1]))
				finB = controlOverflow(int64(uint8(((float64(localTheme[6][2] - localTheme[7][2])) * relB))), int64(localTheme[7][2]))
				structRGBA.Rgba.Set(i, j, color.RGBA{finR, finG, finB, 255})
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}

func Engine2(intensity uint8, structRGBA *structs.StructRGBA, rgbValues []uint8, localTheme [][]int) {

	// hue shift
	rgbR := float64(0)
	rgbG := float64(0)
	rgbB := float64(0)

	var wg sync.WaitGroup

	for i := 0; i < structRGBA.MaxX; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < structRGBA.MaxY; j++ {
				r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
				rgbR += float64(r)
				rgbG += float64(g)
				rgbB += float64(b)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	values := structRGBA.MaxX * structRGBA.MaxY
	rgbR /= float64(values)
	rgbG /= float64(values)
	rgbB /= float64(values)

	diffR := int8(rgbValues[0]) - int8(uint8(rgbR))
	diffG := int8(rgbValues[1]) - int8(uint8(rgbG))
	diffB := int8(rgbValues[2]) - int8(uint8(rgbB))

	for i := 0; i < structRGBA.MaxX; i++ {
		wg.Add(1)
		go func(i int) {
			for j := 0; j < structRGBA.MaxY; j++ {
				r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
				r8 := controlOverflow(int64(r), int64(diffR))
				g8 := controlOverflow(int64(g), int64(diffG))
				b8 := controlOverflow(int64(b), int64(diffB))
				structRGBA.Rgba.Set(i, j, color.RGBA{r8, g8, b8, 255})
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return
}

func Engine3(intensity uint8, structRGBA *structs.StructRGBA, rgbValues []uint8, localTheme [][]int) {
	// match matching colors
	for i := 0; i < structRGBA.MaxX; i++ {
		for j := 0; j < structRGBA.MaxY; j++ {
			r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
			mag := 9999999
			var rM, gM, bM uint8
			for _, val := range localTheme {
				tempMag := abs(val[0] - int(r))
				tempMag += abs(val[1] - int(g))
				tempMag += abs(val[2] - int(b))
				if tempMag < mag {
					mag = tempMag
					rM = uint8(val[0])
					gM = uint8(val[1])
					bM = uint8(val[2])
				}
			}
			var r8, g8, b8 uint8
			if intensity == 0 {
				r8 = controlOverflow(int64(r), int64((float64(rM)-float64(r))/50))
				g8 = controlOverflow(int64(g), int64((float64(gM)-float64(g))/50))
				b8 = controlOverflow(int64(b), int64((float64(bM)-float64(b))/50))
			} else {
				r8 = controlOverflow(int64(r), int64((float64(rM)-float64(r))/float64(intensity)))
				g8 = controlOverflow(int64(g), int64((float64(gM)-float64(g))/float64(intensity)))
				b8 = controlOverflow(int64(b), int64((float64(bM)-float64(b))/float64(intensity)))
			}
			structRGBA.Rgba.Set(i, j, color.RGBA{r8, g8, b8, 255})
		}
	}
	return
}

func AvgValues(structRGBA *structs.StructRGBA) {

	rgbR := float64(0)
	rgbG := float64(0)
	rgbB := float64(0)
	// calculate the avg HSV values of the image
	for i := 0; i < structRGBA.MaxX; i++ {
		for j := 0; j < structRGBA.MaxY; j++ {
			r, g, b, _ := toUint8(structRGBA.Rgba.At(i, j).RGBA())
			rgbR += float64(r)
			rgbG += float64(g)
			rgbB += float64(b)
		}
	}

	values := structRGBA.MaxX * structRGBA.MaxY
	rgbR /= float64(values)
	rgbG /= float64(values)
	rgbB /= float64(values)

	fmt.Printf("rgb %.2f %.2f %.2f\n", rgbR, rgbG, rgbB)
}
