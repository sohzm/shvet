package arguments

import (
	"strconv"
)

var argsMap = map[string]string{
	"-e":          "engine",
	"--engine":    "engine",
	"engine":      "engine",
	"-f":          "force",
	"--force":     "force",
	"force":       "force",
	"-h":          "help",
	"--help":      "help",
	"help":        "help",
	"-i":          "intensity",
	"--intensity": "intensity",
	"intensity":   "intensity",
	"-l":          "list",
	"--list":      "list",
	"list":        "list",
	"-o":          "output",
	"--output":    "output",
	"output":      "output",
	"-v":          "version",
	"--version":   "version",
	"version":     "version",
	"--verbose":   "verbose",
	"verbose":     "verbose",

	"dracula":      "dracula",
	"--dracula":    "dracula",
	"gruvbox":      "gruvbox",
	"--gruvbox":    "gruvbox",
	"nord":         "nord",
	"--nord":       "nord",
	"solarized":    "solarized",
	"--solarized":  "solarized",
	"tokyonight":   "tokyonight",
	"--tokyonight": "tokyonight",
}

var themesAccepted = map[string]int8{
	"nord":       1,
	"gruvbox":    1,
	"tokyonight": 1,
	"dracula":    1,
	"solarized":  1,
}

var optionsAccepted = map[string]int8{
	"help":    1,
	"list":    1,
	"version": 1,
}

func parse(arg string) string {
	temp, ok := argsMap[arg]
	if ok {
		return temp
	} else {
		if _, err := strconv.Atoi(arg); err == nil {
			return arg
		} else {
			return "__checkFileExists__"
		}
	}
}
