package arguments

import (
	"errors"
	"os"
	"shvet/structs"
	"strconv"
)

func ParseArguments(argsArr []string) (argsStruct structs.Args, err error) {
	for i, val := range argsArr {
		// TODO can be improved?
		if parse(val) == "__checkFileExists__" {
			if _, err := os.Stat(val); errors.Is(err, os.ErrNotExist) {
				return argsStruct, errors.New("25b4: Invalid Argument")
			} else {
				argsStruct.Input = val
			}
		}

		if parse(val) == "help" {
			argsStruct.Opt.Help = true
			return
		}

		if parse(val) == "version" {
			argsStruct.Opt.Version = true
			return
		}

		if parse(val) == "list" {
			argsStruct.Opt.List = true
			return
		}

		if parse(val) == "force" {
			argsStruct.Force = true
		}

		if parse(val) == "verbose" {
			argsStruct.Verbose = true
		}

		if parse(val) == "engine" {
			if len(argsArr) <= i+1 {
				return argsStruct, errors.New("f2ec: Invalid Argument")
			}

			tempStr := argsArr[i+1]

			for _, val1 := range tempStr {
				tempStruct := structs.Intensities{
					Engine: string(val1),
				}

				if val1 == '1' {
					tempStruct.Intensity = 50
				} else if val1 == '2' {
					tempStruct.Intensity = 50
				} else if val1 == '3' {
					tempStruct.Intensity = 50
				} else {
					return argsStruct, errors.New("0d16: Invalid Argument")
				}

				argsStruct.EngineArr = append(argsStruct.EngineArr, tempStruct)
			}
		}

		// intensity should always come after engine
		if parse(val) == "intensity" {
			if len(argsStruct.EngineArr) == 0 {
				return argsStruct, errors.New("6043: Invalid Argument")
			}
			if len(argsArr) <= i+len(argsStruct.EngineArr) {
				return argsStruct, errors.New("40b9: Invalid Argument")
			}
			for ele := range argsStruct.EngineArr {
				if tempInt, err := strconv.Atoi(argsArr[ele+i+1]); err == nil {
					argsStruct.EngineArr[ele].Intensity = tempInt
				} else {
					return argsStruct, errors.New("2c35: Invalid Argument")
				}
			}
		}
	}

	// TODO check if the struct is empty
	return
}
