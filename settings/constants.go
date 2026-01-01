package settings

const (
	WorldWidth  = 800
	WorldHeight = 800
	UIWidth     = 240
	UIHeight    = WorldHeight

	ScreenHeight = WorldHeight
	ScreenWidth  = WorldWidth + UIWidth

	DisplayScale = 1

	ColourSaturation = 1.0
	ColourValue      = 0.6
)

var (
	MaxInfluenceRadius       float64 = 50
	UniversalForceMultiplier float64 = 0.6
	Friction                 float64 = 0.92
	CursorRepelForce         float64 = 0.015

	InverseMaxInfluenceRadiusSquared float64 = 1 / (MaxInfluenceRadius * MaxInfluenceRadius)
)
