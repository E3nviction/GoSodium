package sodium

import (
	"fmt"
	"sodium/components"
	"sodium/types"
)

func DrawUI(uiRoot []types.Element, window *types.Window) {
	iterateUI(uiRoot, window, nil)
}

func PushUI(filename string, window *types.Window) {
	UI, err := LoadUI(filename)
	if err != nil {
		fmt.Println("Error loading UI:", err)
		return
	}
	iterateUI(UI, window, nil)
}

// iterateUI recursively processes each element in the UI structure
func iterateUI(uiRoot []types.Element, window *types.Window, parent *Box) {
	for _, child := range uiRoot {
		if child.Tag == "rect" {
			components.HandleRect(child, window)
		}
		if child.Tag == "button" {
			components.HandleButton(child, window)
		}
		iterateUI(child.Children, window, parent)
	}
}
