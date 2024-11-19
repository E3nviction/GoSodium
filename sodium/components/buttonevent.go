package components

import (
	"sodium/types"

	"github.com/veandco/go-sdl2/sdl"
)

var buttonState = false

// HandleButtonEvent adds a button press-and-release handler
func HandleButtonClick(win *types.Window, x, y, width, height int, handler func(element types.Element), self types.Element) {
	mouseX, mouseY, mouseState := sdl.GetMouseState()

	// Check if the mouse is within the button's rectangle
	insideButton := int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height

	// Check if the left mouse button is pressed
	isPressed := (mouseState & sdl.Button(sdl.BUTTON_LEFT)) != 0

	// If the button is pressed and the mouse is inside the button, set buttonState to true
	if isPressed && insideButton {
		buttonState = true
	}

	// If the button is released
	if !isPressed {
		// Only call the handler if buttonState is true and the mouse is inside the button
		if buttonState && insideButton {
			handler(self)
		}

		// Reset buttonState on mouse release
		buttonState = false
	}
}

func HandleButtonPress(win *types.Window, x, y, width, height int, pressHandler func(element types.Element), self types.Element) {
	mouseX, mouseY, mouseState := sdl.GetMouseState()

	// Check if the mouse is within the button's rectangle
	insideButton := int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height

	// Check if the left mouse button is pressed
	isPressed := (mouseState & sdl.Button(sdl.BUTTON_LEFT)) != 0

	// If the button is pressed and the mouse is inside the button, call the press handler
	if isPressed && insideButton {
		pressHandler(self)
	}
}

func HandleButtonHover(win *types.Window, x, y, width, height int, hoverHandler func(element types.Element), self types.Element) {
	mouseX, mouseY, _ := sdl.GetMouseState()

	// Check if the mouse is within the button's rectangle
	insideButton := int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height

	// Call hover handler if the mouse is inside the button
	if insideButton {
		hoverHandler(self)
	}
}

func HandleButtonUnFocus(win *types.Window, x, y, width, height int, unfocusHandler func(element types.Element), self types.Element) {
	mouseX, mouseY, _ := sdl.GetMouseState()
	insideButton := int(mouseX) >= x && int(mouseX) <= x+width &&
		int(mouseY) >= y && int(mouseY) <= y+height

	// Call unfocus handler if the mouse is not inside the button
	if !insideButton {
		unfocusHandler(self)
	}
}
