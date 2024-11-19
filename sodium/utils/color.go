package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func ConvertColor(colorStr string, alphaStr string) (sdl.Color, error) {
	var color uint32
	if strings.HasPrefix(colorStr, "#") {
		colorStr = colorStr[1:] // Remove the "#"
		if len(colorStr) != 6 {
			return sdl.Color{}, fmt.Errorf("color must be in the format #RRGGBB")
		}

		_, err := fmt.Sscanf(colorStr, "%x", &color)
		if err != nil {
			return sdl.Color{}, fmt.Errorf("invalid color: %s", colorStr)
		}
	} else {
		return sdl.Color{}, fmt.Errorf("invalid color format. Use #RRGGBB.")
	}

	r := uint8(color >> 16)
	g := uint8(color >> 8 & 0xFF)
	b := uint8(color & 0xFF)

	alpha := uint8(255) // Default to fully opaque
	if alphaStr != "" {
		alphaValue, err := strconv.Atoi(alphaStr)
		if err == nil {
			alpha = uint8(alphaValue)
		}
	}

	return sdl.Color{R: r, G: g, B: b, A: alpha}, nil
}

func Color(r, g, b, a uint8) sdl.Color {
	return sdl.Color{R: r, G: g, B: b, A: a}
}
