package arguments

import (
	"errors"
	"os"
	"shvet/data"
	"shvet/structs"
)

func ParseArguments(argsArr []string) (argsStruct structs.Args, err error) {
	for _, val := range argsArr {
		if parse(val) == "help" {
			argsStruct.OptionBool = true
			argsStruct.Opt.Help = true
			return
		}

		if parse(val) == "version" {
			argsStruct.OptionBool = true
			argsStruct.Opt.Version = true
			return
		}

		if parse(val) == "list" {
			argsStruct.OptionBool = true
			argsStruct.Opt.List = true
			return
		}

		if _, ok := data.DataMap[parse(val)]; ok {
			argsStruct.Theme = parse(val)
			continue
		}

		// TODO can be improved?
		if parse(val) == "__checkFileExists__" {
			if _, err := os.Stat(val); errors.Is(err, os.ErrNotExist) {
				return argsStruct, errors.New("25b4: " + err.Error())
			} else {
				argsStruct.Input = val
			}
		}

		if parse(val) == "force" {
			argsStruct.Force = true
		}

		if parse(val) == "verbose" {
			argsStruct.Verbose = true
		}

		if parse(val) == "customize" {
			argsStruct.Customize = true
			return argsStruct, errors.New("39ad: Customize isn't added, yet. Contributions are welcome at https://github.com/sz47/shvet")
		}
	}

	if argsStruct.Theme == "" {
		return argsStruct, errors.New("74b5: Oops, you forgot theme, \nor maybe you spelled it wrong?\nYou can use `shvet list` for listing theme names")
	}
	if argsStruct.Input == "" {
		return argsStruct, errors.New("932d: Oops, you forgot input file, \nor maybe you spelled it wrong?")
	}
	return
}
