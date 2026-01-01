package main

import (
	"image/color"
	"life/particle"
	"life/settings"
	"life/ui"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	NParticles = 100
	NVariants  = 6
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

	dt := 1.0 / float64(ebiten.TPS())

	g.particles.Update(dt)
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.particles.Draw(screen)
	g.ui.Draw(screen)
	g.ui.DrawDebugUI(screen)

	ebitenutil.DebugPrintAt(
		screen,
		"FPS: "+strconv.Itoa(int(ebiten.ActualFPS()))+" TPS: "+strconv.Itoa(int(ebiten.ActualTPS())),
		settings.WorldWidth+10,
		10,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.ScreenWidth, settings.ScreenHeight
}

func main() {
	s := ebiten.Monitor().DeviceScaleFactor()
	ebiten.SetWindowSize(
		int(settings.ScreenWidth*settings.DisplayScale/s),
		int(settings.ScreenHeight*settings.DisplayScale/s),
	)
	ebiten.SetWindowTitle("Particle Life")

	g := &Game{}
	g.particles = particle.NewRandomParticleSet(NParticles, NVariants)
	g.ui = ui.NewUI(g.particles)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
