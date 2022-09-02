package handlers

import (
	"shvet/converters"
	"shvet/files"
	"shvet/operations"
	"shvet/structs"
)

func Handle(args structs.Args) {
	if args.Opt.Version {
		handleVersion()
		return
	}
	if args.Opt.Help {
		handleHelp()
		return
	}
	if args.Opt.List {
		handleList()
		return
	}
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
