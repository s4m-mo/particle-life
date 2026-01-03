package particle

import (
	"image/color"
	"life/settings"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	x  float64
	y  float64
	vx float64
	vy float64

	variant uint
}

// Update the position of the particle according to its current velocity
func (p *Particle) Update() {
	p.x = PositiveModulus(p.x+p.vx, settings.WorldWidth)
	p.y = PositiveModulus(p.y+p.vy, settings.WorldHeight)

	p.vx *= settings.Friction
	p.vy *= settings.Friction
}

// This is necessary as math.Mod returns a negative value for negative inputs
func PositiveModulus(v float64, m float64) float64 {
	for v < 0 {
		v += m
	}

	return math.Mod(v, m)
}

// Draw the particle to the screen (set a single pixel to be used by the shader later)
func (p *Particle) Draw(screen *ebiten.Image, col color.Color) {
	screen.Set(int(p.x), int(p.y), col)
}

/* Get the quadrants around the particle in which other particles can influence it. */
func (p *Particle) FindInfluencingQuadrants() [4][2]int {
	// (nx, ny) is the quadrant that the centre of the particle is in
	nx, ny := p.CurrentQuadrant()

	relx := math.Mod(p.x, settings.MaxInfluenceRadius) / settings.MaxInfluenceRadius
	rely := math.Mod(p.y, settings.MaxInfluenceRadius) / settings.MaxInfluenceRadius

	var offsetX, offsetY int
	if relx > 0.5 {
		offsetX = 1
	} else {
		offsetX = -1
	}

	if rely > 0.5 {
		offsetY = 1
	} else {
		offsetY = -1
	}

	return [4][2]int{
		{nx, ny},
		{nx + offsetX, ny},
		{nx, ny + offsetY},
		{nx + offsetX, ny + offsetY},
	}
}

/* Get the quadrant indices of the particle. */
func (p *Particle) CurrentQuadrant() (int, int) {
	nx := int(math.Floor(p.x / settings.MaxInfluenceRadius)) // Truncates as required automatically
	ny := int(math.Floor(p.y / settings.MaxInfluenceRadius))

	return nx, ny
}

// Info about a particle in a quadrant
type QuadInfo struct {
	X int
	Y int
	P *Particle
}
