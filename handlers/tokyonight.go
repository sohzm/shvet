package handlers

import (
	"shvet/converters"
	"shvet/data"
	"shvet/files"
	"shvet/operations"
	"strconv"
)

func HandleTokyonight(fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	operations.Engine(&structRGBA, data.GruvboxPoints)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

func HandleTokyonight2(intensity, fileName string) error {
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

	operations.Engine2(uint8(intensityInt), &structRGBA, data.TokyonightRGB, data.TokyonightColors)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}
