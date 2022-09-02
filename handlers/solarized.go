package handlers

import (
	"shvet/converters"
	"shvet/data"
	"shvet/files"
	"shvet/operations"
	"strconv"
)

func HandleSolarized(fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	operations.Engine(&structRGBA, data.SolarizedPoints)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

func HandleSolarized2(intensity, fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	intensityInt, err := strconv.Atoi(intensity)
	if err != nil {
		return err
	}

	operations.Engine2(uint8(intensityInt), &structRGBA, data.SolarizedRGB, data.SolarizedColors)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}
