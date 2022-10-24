package structs

import "image"

type StructRGBA struct {
	Rgba       *image.RGBA
	MaxX, MaxY int
}

type Args struct {
	Verbose    bool
	Force      bool
	Theme      string
	Input      string
	Output     string
	OptionBool bool
	Opt        Opt
	Customize  bool
	EngineArr  []Intensities
}

type Opt struct {
	Version bool
	Help    bool
	List    bool
}

type Intensities struct {
	Engine    string
	Intensity int
}

type Color struct {
	R, G, B uint8
}
