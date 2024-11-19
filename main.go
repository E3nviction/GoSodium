package main

import (
	"fmt"
	"sodium"
	"sodium/types"
)

func Init(w *types.Window) {
	UI, _ := sodium.LoadUI("main.ui")
	w.AddtoDictionary("UI", UI)
	w.Handlers = map[string]func(element types.Element){
		"Click": func(self types.Element) {
			fmt.Println("Hello, World!")
		},
		"Offfocus": func(self types.Element) {
			self.Attributes["BackgroundColor"] = "#222222"
		},
		"Press": func(self types.Element) {
			self.Attributes["BackgroundColor"] = "#ff2222"
		},
		"Hover": func(self types.Element) {
			self.Attributes["BackgroundColor"] = "#552222"
		},
	}
	sodium.DrawWindow(UI, w)
}

func Loop(w *types.Window) {
	UI := w.GetFromDictionary("UI")
	sodium.DrawUI(UI.([]types.Element), w)
}

func main() {
	sodium.Start(Init, Loop)
}
