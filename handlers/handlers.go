package handlers

import (
	"fmt"
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

func HandleList() {
	fmt.Println("Themes supported in Shvet:")
	fmt.Println("   - dracula")
	fmt.Println("   - gruvbox")
	fmt.Println("   - nord")
	fmt.Println("   - solarized")
	fmt.Println("   - tokyonight")
}

func HandleHelp() {
	fmt.Println("Shvet")
	fmt.Println(" - A color scheme adjuster for your wallpapers")
	fmt.Println("")
	fmt.Println("more info at https://github.com/sz47/shvet")
}
