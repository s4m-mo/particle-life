package ui

import (
	"life/particle"

	"github.com/hajimehoshi/ebiten/v2"
)

type UI struct {
	particles *particle.ParticleSet
}

func NewUI(particles *particle.ParticleSet) *UI {
	return &UI{
		particles: particles,
	}
}

func (ui *UI) Update() {
}

func (ui *UI) Draw(screen *ebiten.Image) {
}
