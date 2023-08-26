package structs

import "image"

type StructRGBA struct {
	Rgba       *image.RGBA
	MaxX, MaxY int
}

type Args struct {
	Verbose   bool
	Force     bool
	Theme     string
	Input     string
	Output    string
	BlendMode string
	Opt       Opt
	//Customize  bool
	//EngineArr  []Intensities
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

type Colorx struct {
	RGB RGBx
	HSV HSVx
}

type RGBx struct {
	R, G, B uint8
}

type HSVx struct {
	H uint
	S uint8
	V uint8
}

type Tone struct {
}
