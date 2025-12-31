package main

import (
	"image/color"
	"life/particle"
	"life/settings"
	"life/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	NParticles = 900
	NVariants  = 10
)

type Game struct {
	particles *particle.ParticleSet
	ui        *ui.UI
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.particles.RegenerateAttractionMatrix()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.particles.RegenerateCentralPoints(NParticles, NVariants)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.ui.ToggleDebugUI()
	}

	g.particles.Update()
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.particles.Draw(screen)
	g.ui.Draw(screen)
	g.ui.DrawDebugUI(screen)
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
	g.particles = particle.NewCentredParticleSet(NParticles, NVariants)
	g.ui = ui.NewUI(g.particles)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
