package utils

import (
	"image/color"
	"math"
)

func HSVToRGB(h, s, v float64) color.Color {
	var R, G, B float64
	var C, X, m float64

	C = v * s
	X = C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m = v - C

	h = math.Mod(h, 360) // Ensure h is within 0-359

	switch {
	case h < 60:
		R, G, B = C, X, 0
	case h < 120:
		R, G, B = X, C, 0
	case h < 180:
		R, G, B = 0, C, X
	case h < 240:
		R, G, B = 0, X, C
	case h < 300:
		R, G, B = X, 0, C
	case h < 360:
		R, G, B = C, 0, X
	default:
		return color.White
	}

	return color.RGBA{
		uint8((R + m) * 255),
		uint8((G + m) * 255),
		uint8((B + m) * 255),
		255,
	}
}
