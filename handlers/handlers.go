package handlers

import (
	"shvet/converters"
	"shvet/data"
	"shvet/files"
	"shvet/operations"
	"shvet/structs"
)

func Handle(args structs.Args) error {
	if args.OptionBool {
		if args.Opt.Version {
			handleVersion()
			return nil
		}
		if args.Opt.Help {
			handleHelp()
			return nil
		}
		if args.Opt.List {
			handleList()
			return nil
		}
	}

	return handleTheme(args)
}

func handleTheme(args structs.Args) error {
	img, err := files.Open(args.Input)
	if err != nil {
		return err
	}

	structRGBA, err := converters.ImageToRGBA(img)
	if err != nil {
		return err
	}

	if args.Force {
		operations.Engine3(1, &structRGBA, data.DataMap[args.Theme].RGB, data.DataMap[args.Theme].Colors)
	} else {
		operations.Engine(&structRGBA, data.DataMap[args.Theme].Points)
	}
	err = files.Write(args.Input, structRGBA.Rgba)
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
