package sodium

import (
	"fmt"
	"sodium/types"
	"strconv"
)

func DrawWindow(uiRoot []types.Element, window *types.Window) {
	iterateWindow(uiRoot, window)
}

func PushWindow(filename string, window *types.Window) {
	UI, err := LoadUI(filename)
	if err != nil {
		fmt.Println("Error loading UI:", err)
		return
	}
	iterateWindow(UI, window)
}

// iterateUI recursively processes each element in the UI structure to locate "window" elements
func iterateWindow(uiRoot []types.Element, window *types.Window) {
	for _, child := range uiRoot {
		if child.Tag == "window" {
			width := child.Attributes["Width"]
			height := child.Attributes["Height"]

			// Convert string to int32
			widthInt, err := strconv.Atoi(width)
			if err != nil {
				fmt.Println("Invalid width:", width)
				return
			}
			heightInt, err := strconv.Atoi(height)
			if err != nil {
				fmt.Println("Invalid height:", height)
				return
			}

			// Validate width and height attributes.
			if width == "" || height == "" {
				fmt.Println("Width and Height attributes are required for the window.")
				return
			}
			window.Window.SetSize(int32(widthInt), int32(heightInt))
			_, err = window.Window.GetSurface()
			if err != nil {
				fmt.Println("Error retrieving new window surface:", err)
				return
			}

			window.Window.UpdateSurface()
		}
		iterateWindow(child.Children, window)
	}
}
