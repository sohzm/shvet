package handlers

import (
	"github.com/sz47/shvet/converters"
	"github.com/sz47/shvet/files"
	"github.com/sz47/shvet/operations"
	"github.com/sz47/shvet/structs"
)

func Handle(parsedArgs structs.Args) error {

	switch {
	case parsedArgs.Opt.Version:
		handleVersion()
		return nil
	case parsedArgs.Opt.Help:
		handleHelp()
		return nil
	case parsedArgs.Opt.List:
		handleList()
		return nil
	}

	img, err := files.Open(parsedArgs.Input)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	if parsedArgs.Force {
		//operations.Engine3(1, &structRGBA, data.DataMap[args.Theme].SmallPalatte)
	} else {
		operations.Engine4(&structRGBA, parsedArgs.Theme, parsedArgs.BlendMode)
	}
	err = files.Write(parsedArgs.Input, structRGBA.Rgba)
	if err != nil {
		return err
	}

	return nil
}

// TODO see if this is needed
//func handleTheme2(args structs.Args) error {
//	img, err := files.Open(args.Input)
//	if err != nil {
//		return err
//	}
//
//	structRGBA, err := converters.ImageToRGBA(img)
//	if err != nil {
//		return err
//	}
//
//	//intensityInt, err := strconv.Atoi(intensity)
//	if err != nil {
//		return err
//	}
//
//	operations.Engine2(uint8(intensityInt), &structRGBA, data.GruvboxRGB, data.GruvboxColors)
//	err = files.Write(fileName, structRGBA.Rgba)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
