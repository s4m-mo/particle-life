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

func (p *Particle) Update() {
	p.x = math.Mod((p.x + p.vx), settings.WorldWidth)
	p.y = math.Mod((p.y + p.vy), settings.WorldHeight)
}

func (p *Particle) Draw(screen *ebiten.Image, col color.Color) {
	screen.Set(int(p.x), int(p.y), col)
}
