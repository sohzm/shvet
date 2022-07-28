package handlers

import (
	"shvet/converters"
	"shvet/files"
	"shvet/operations"
	"strconv"
)

func HandleDarken(intensity, fileName string) error {
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

	operations.Darken(uint8(intensityInt), &structRGBA)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

func HandleLighten(intensity, fileName string) error {
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

	operations.Lighten(uint8(intensityInt), &structRGBA)
	err = files.Write(fileName, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

func HandleAvgValues(fileName string) error {
	img, err := files.Open(fileName)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	operations.AvgValues(&structRGBA)
	err = files.Write(fileName, structRGBA.Rgba)
	return nil
}
