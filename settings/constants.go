package settings

const (
	WorldWidth  = 1200
	WorldHeight = 1000
	UIWidth     = 360
	UIHeight    = WorldHeight

	UIPadX           = 20
	UIPadY           = 10
	UIColourBarWidth = 10

	MinWorldDim            = min(WorldHeight, WorldWidth)
	HalfMinWorldDimSquared = (MinWorldDim * MinWorldDim) / 2

	ScreenHeight = WorldHeight
	ScreenWidth  = WorldWidth + UIWidth

	DisplayScale = 1

	ColourSaturation = 0.9
	ColourValue      = 0.75

	CursorRadiusChangeSpeed = 1.04 * 60 // To account for dt at 60fps
)

var (
	MaxInfluenceRadius       float64 = 60
	UniversalForceMultiplier float64 = 0.6
	Friction                 float64 = 0.924
	CursorRepelForce         float64 = 0.015
	CursorForceRadiusSquared float64 = 625.0 // (radius squared to save computation on the square root)

	InverseMaxInfluenceRadiusSquared float64 = 1 / (MaxInfluenceRadius * MaxInfluenceRadius)
)
