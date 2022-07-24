package files

import (
    "os"

    "io/ioutil"
    "image"
    "github.com/h2non/filetype"
	"image/jpeg"
    "image/png"
    "errors"
)

func Open (fileName string) (image.Image, error) {
    file, err := os.Open(fileName)
    if (err != nil) {
        return nil, err
    }

    buf, _ := ioutil.ReadFile(fileName)
    kind, _ := filetype.Match(buf)
    if kind == filetype.Unknown {
        return nil, errors.New("Filetype Unknown")
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
