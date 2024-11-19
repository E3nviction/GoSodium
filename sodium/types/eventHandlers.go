package types

import (
	"github.com/veandco/go-sdl2/sdl"
)

func HandleMouseEvent(win *Window, x, y, width, height int, handler func()) {
	mouseX, mouseY, mouseState := sdl.GetMouseState()

	// Check if the mouse is within the rectangle and a click occurred
	if int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height &&
		(mouseState&sdl.Button(sdl.BUTTON_LEFT)) != 0 {

		// Add the event to the event map
		AddEvent("MouseClick", handler)
	}

	// Check if the mouse is within the rectangle
	if int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height {
		// Add the event to the event map
		AddEvent("MouseHover", handler)
	}

	// Check if the mouse is not within the rectangle
	if int(mouseX) < x || int(mouseX) > x+width ||
		int(mouseY) < y || int(mouseY) > y+height {
		// Add the event to the event map
		AddEvent("MouseUnFocus", handler)
	}

	// Check if the mouse is pressed
	if (mouseState & sdl.Button(sdl.BUTTON_LEFT)) != 0 {
		// Add the event to the event map
		AddEvent("MousePress", handler)
	}
}
