package data

type PalatteData struct {
	AverageRGB   []uint8
	SmallPalatte [][]uint8
	LargePalatte [][]uint8
}

var ImpPoints = [][]uint8{
	{255, 0, 0},     // red
	{255, 255, 0},   // ylo
	{0, 255, 0},     // green
	{0, 255, 255},   // cyan
	{0, 0, 255},     // bluwu
	{255, 0, 255},   // mag
	{255, 255, 255}, // white
	{0, 0, 0},       // black
}

var DataMap = map[string]PalatteData{
	"nord":       Nord,
	"tokyonight": Tokyonight,
	"solarized":  Solarized,
	"dracula":    Dracula,
	"gruvbox":    Gruvbox,
}
