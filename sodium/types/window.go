package types

import "github.com/veandco/go-sdl2/sdl"

type Window struct {
	Components []Component
	Window     *sdl.Window
	Renderer   *sdl.Renderer
	Handlers   map[string]func(element Element)
	Dictionary map[string]interface{}
}

func (w *Window) Add(component Component) {
	w.Components = append(w.Components, component)
}

func (w *Window) ShowAll() {
	for _, c := range w.Components {
		c.Draw(w.Window)
	}
}

func (w *Window) FillSurface(color uint32) {
	_, err := w.Window.GetSurface()
	if err != nil {
		panic(err)
	}
}

func (w *Window) UpdateSurface() error {
	return w.Window.UpdateSurface()
}

func (w *Window) GetSurface() (*sdl.Surface, error) {
	return w.Window.GetSurface()
}

func (s *Window) AddtoDictionary(key string, value interface{}) {
	s.Dictionary[key] = value
}

func (s *Window) GetFromDictionary(key string) interface{} {
	return s.Dictionary[key]
}

func (s *Window) RemoveFromDictionary(key string) {
	delete(s.Dictionary, key)
}

func (s *Window) ClearDictionary() {
	s.Dictionary = map[string]interface{}{}
}
