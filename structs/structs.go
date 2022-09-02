package structs

import "image"

type StructRGBA struct {
	Rgba       *image.RGBA
	MaxX, MaxY int
}

type Args struct {
	EngineArr []Intensities
	Verbose   bool
	Force     bool
	Opt       Opt
	Input     string
	Output    string
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
