package sodium

import "github.com/veandco/go-sdl2/sdl"

type Box struct {
	X, Y, Width, Height int
	Margin, Padding     int
	Color               sdl.Color
	Rounding            string
}
