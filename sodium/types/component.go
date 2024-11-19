package types

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	Draw(window *sdl.Window)
}
