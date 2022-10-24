package handlers

import (
	"fmt"
	"shvet/converters"
	"shvet/files"
	"shvet/operations"
)

func handleList() {
	fmt.Println("Themes supported in Shvet:")
	fmt.Println("   - dracula")
	fmt.Println("   - gruvbox")
	fmt.Println("   - nord")
	fmt.Println("   - solarized")
	fmt.Println("   - tokyonight")
}

func handleHelp() {
	fmt.Println("Shvet")
	fmt.Println(" - A color scheme adjuster for your wallpapers")
	fmt.Println("")
	fmt.Println("more info at https://github.com/sz47/shvet")
}

func handleVersion() {
	fmt.Println("Shvet, version 0.7")
}

func handleAvgValues(fileName string) error {
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
