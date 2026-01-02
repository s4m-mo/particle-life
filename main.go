package main

import (
	"fmt"
	"image/color"
	"life/particle"
	"life/settings"
	"life/ui"
	"life/ui/elems"
	"life/utils"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	NParticles = 2000
	NVariants  = 4
)

var (
	FourthRootOfTwo = math.Pow(2, 0.25)
)

type Game struct {
	particles *particle.ParticleSet
	ui        *ui.UI
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		g.ui.ToggleDebugUI()
	}

	if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		settings.CursorForceRadiusSquared = utils.Clamp(
			(settings.CursorForceRadiusSquared / (settings.CursorRadiusChangeSpeed * dt)),
			100,
			settings.HalfMinWorldDimSquared*FourthRootOfTwo,
		)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
		settings.CursorForceRadiusSquared = utils.Clamp(
			(settings.CursorForceRadiusSquared * (settings.CursorRadiusChangeSpeed * dt)),
			100,
			settings.HalfMinWorldDimSquared*FourthRootOfTwo,
		)
	}

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
		settings.WorldWidth+settings.UIPadX,
		settings.WorldHeight-20,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.ScreenWidth, settings.ScreenHeight
}

func SetupUI(u *ui.UI, p *particle.ParticleSet) {
	const offsetX = settings.WorldWidth + settings.UIPadX

	T1 := color.RGBA{30, 30, 50, 255}
	T2 := color.RGBA{40, 35, 60, 255}
	T3 := color.RGBA{20, 20, 35, 255}

	W := settings.UIWidth - 2*settings.UIPadX
	N := 0
	H := settings.UIHeight - 2*settings.UIPadY

	Cx := W / 12
	Cy := H / 48
	// Nh := H / Cx

	titleReset := elems.NewTextElement("Reset:", offsetX, (2*Cx-elems.FontLineHeight)/2+settings.UIPadY)
	u.AddElement(titleReset)

	btnWidth := (settings.UIWidth - (2 * settings.UIPadX) - titleReset.Width()) / 3

	u.AddElement(
		elems.NewButton(
			"1",
			settings.WorldWidth+titleReset.Width()+settings.UIPadX, settings.UIPadY,
			btnWidth, Cx*2,
			func() { p.RegenerateUniformPoints(NParticles, NVariants) },
			T1, T2, T3,
		),
	)

	u.AddElement(
		elems.NewButton(
			"2",
			settings.WorldWidth+titleReset.Width()+settings.UIPadX+btnWidth, settings.UIPadY,
			btnWidth, Cx*2,
			func() { p.RegenerateCentralPoints(NParticles, NVariants) },
			T1, T2, T3,
		),
	)

	u.AddElement(
		elems.NewButton(
			"3",
			settings.WorldWidth+titleReset.Width()+settings.UIPadX+2*btnWidth, settings.UIPadY,
			btnWidth, Cx*2,
			func() { p.RegenerateOuterPoints(NParticles, NVariants) },
			T1, T2, T3,
		),
	)

	N += 3

	u.AddElement(elems.NewTextElement("Attraction Matrix", offsetX, (N*Cx)+settings.UIPadY))

	u.AddElement(
		elems.NewButton(
			"Random Matrix",
			offsetX, (N+1)*Cx+settings.UIPadY,
			Cx*12, Cx,
			func() { p.RegenerateAttractionMatrix() },
			T1, T2, T3,
		),
	)

	N += 3

	// scrollNotice := elems.NewTextElement("Use mouse scrollwheel to adjust values", offsetX, (N*Cx)+settings.UIPadY+3)
	// scrollNotice.Size = 11
	// u.AddElement(scrollNotice)

	cols := p.GetVariantColours()
	const colourBarRemaining = settings.UIWidth - settings.UIColourBarWidth - 2*settings.UIPadX

	variantCellSize := (colourBarRemaining / len(cols))
	for i, v := range cols {
		u.AddElement(
			elems.NewBox(
				offsetX+settings.UIColourBarWidth+variantCellSize*i,
				(N*Cx)+settings.UIPadY,
				colourBarRemaining/len(cols),
				settings.UIColourBarWidth,
				v,
			),
		)

		u.AddElement(
			elems.NewBox(
				offsetX,
				(N*Cx)+settings.UIPadY+settings.UIColourBarWidth+variantCellSize*i,
				settings.UIColourBarWidth,
				colourBarRemaining/len(cols),
				v,
			),
		)

		for j, _ := range cols {
			u.AddElement(
				elems.NewScrollMeter(
					offsetX+settings.UIColourBarWidth+(variantCellSize*i),
					N*Cx+settings.UIPadY+settings.UIColourBarWidth+(variantCellSize*j),
					variantCellSize, variantCellSize,
					func(initial, wheel float64) float64 {
						return initial + (wheel * 0.01)
					},
					&p.AttractionMatrix[j][i],
					-1, 1,
				),
			)
		}
	}

	N += 2 + (len(cols) * variantCellSize / Cx)

	const setBtnWidth = (settings.UIWidth - settings.UIPadX*2) / 3
	for i := 0; i < 3; i++ {
		u.AddElement(
			elems.NewButton(
				fmt.Sprint(i-1),
				offsetX+setBtnWidth*i, N*Cx+settings.UIPadY,
				setBtnWidth, Cx*1,
				func() { p.SetAttractionMatrix(float64(i - 1)) },
				T1, T2, T3,
			),
		)
	}

	helpText := "[L/R Click] Interact\n[-/+] Change cursor size\n[G] Toggle Debug Grid"
	helpTextSections := strings.Split(helpText, "\n")

	for i, t := range helpTextSections {
		el := elems.NewTextElement(
			t,
			offsetX, H-(settings.UIPadY)-((len(helpTextSections)-i)*Cy),
		)

		el.Size = 14
		u.AddElement(el)
	}

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

	SetupUI(g.ui, g.particles)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
