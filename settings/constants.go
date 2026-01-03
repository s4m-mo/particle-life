package settings

const (
	// World Size
	WorldWidth  = 1200
	WorldHeight = 1000

	// UI Settings
	UIWidth          = 360
	UIHeight         = WorldHeight
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
	MaxInfluenceRadius       float64 = 60    // The maximum distance between which a force acts between 2 particles - also determines the grid size of the space partitions
	UniversalForceMultiplier float64 = 0.6   // Multiplies all forces by this value
	Friction                 float64 = 0.924 // Friction particles have (to prevent exploding velocities)
	CursorRepelForce         float64 = 0.015 // Strength of cursor interaction force
	CursorForceRadiusSquared float64 = 625.0 // (radius squared to save computation on the square root)

	InverseMaxInfluenceRadiusSquared float64 = 1 / (MaxInfluenceRadius * MaxInfluenceRadius) // Precomputed value for performance
)
