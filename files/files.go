package files

import (
	"os"

	"errors"
	"github.com/h2non/filetype"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
)

func Open(fileName string) (image.Image, error) {
	// TODO make separate func for open and decode
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	buf, _ := ioutil.ReadFile(fileName)
	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		return nil, errors.New("ea02: Filetype Unknown")
	}

	var img image.Image
	mimeType := kind.MIME.Value
	if mimeType == "image/jpeg" {
		img, _ = jpeg.Decode(file)
	}
	if mimeType == "image/png" {
		img, _ = png.Decode(file)
	}

	return img, nil
}

func Write(fileName string, rgba *image.RGBA) error {
	// TODO make separate func for writing and encoding
	// TODO add custom naming option
	f1, err := os.Create("temp_" + fileName + ".png")
	if err != nil {
		return err
	}
	png.Encode(f1, rgba)

	return nil
}
