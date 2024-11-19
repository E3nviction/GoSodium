package sodium

import (
	"sodium/types"

	"github.com/veandco/go-sdl2/sdl"
)

func Start(init func(w *types.Window), loop func(w *types.Window)) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	NewWindow, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer NewWindow.Destroy()

	Renderer, err := sdl.CreateRenderer(NewWindow, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer Renderer.Destroy()

	win := &types.Window{Window: NewWindow, Renderer: Renderer, Dictionary: map[string]interface{}{}}

	init(win)

	NewWindow.UpdateSurface()

	running := true
	for running {
		// Handle SDL events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// Clear the screen
		win.Renderer.SetDrawColor(0, 0, 0, 255)
		win.Renderer.Clear()

		// User-defined loop logic
		loop(win)

		// Process all custom events from the event map
		for _, handler := range types.EventMap {
			handler()
		}

		// Clear the event map for the next iteration
		types.ClearEvents()

		// Present the renderer
		win.Renderer.Present()

		sdl.Delay(10)
	}
}
