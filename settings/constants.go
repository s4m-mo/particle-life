package settings

const (
	WorldWidth  = 400
	WorldHeight = 400
	UIWidth     = 240
	UIHeight    = WorldHeight

	ScreenHeight = WorldHeight
	ScreenWidth  = WorldWidth + UIWidth

	DisplayScale = 2

	ColourSaturation = 1.0
	ColourValue      = 0.8
)

var (
	MaxInfluenceRadius       float64 = 100
	UniversalForceMultiplier float64 = 0.01
	Friction                 float64 = 0.9

	InverseMaxInfluenceRadiusSquared float64 = 1 / (MaxInfluenceRadius * MaxInfluenceRadius)
)
