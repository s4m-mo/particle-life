package main

import (
	"image/color"
	"life/particle"
	"life/settings"
	"life/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	particles *particle.ParticleSet
	ui        *ui.UI
}

func (g *Game) Update() error {
	g.particles.Update()
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.particles.Draw(screen)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.ScreenWidth, settings.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(
		settings.ScreenWidth*settings.DisplayScale,
		settings.ScreenHeight*settings.DisplayScale,
	)
	ebiten.SetWindowTitle("Particle Life")

	g := &Game{}
	g.particles = particle.NewCentredParticleSet(10, 5)
	g.ui = ui.NewUI(g.particles)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
