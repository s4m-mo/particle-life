package ui

import (
	_ "embed"
	"image"
	"life/particle"
	"life/settings"
	"life/ui/elems"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed gridshader.kage
	gridShaderSource []byte

	shader     *ebiten.Shader
	gridShader *ebiten.Shader
)

func init() {
	var err error

	if gridShader, err = ebiten.NewShader(gridShaderSource); err != nil {
		log.Fatal(err)
	}
}

type UI struct {
	particles *particle.ParticleSet

	wasDown bool
	debugUI bool

	elements []*elems.Element
}

func NewUI(particles *particle.ParticleSet) *UI {
	return &UI{
		particles: particles,
	}
}

func (ui *UI) ToggleDebugUI() {
	ui.debugUI = !ui.debugUI
}

func (ui *UI) AddElement(elem elems.Element) {
	ui.elements = append(ui.elements, &elem)
}

func (ui *UI) Update() {
	cx, cy := ebiten.CursorPosition()
	lmb := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	_, dy := ebiten.Wheel()

	for _, el := range ui.elements {
		(*el).Update(cx, cy, dy, lmb)
	}
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, el := range ui.elements {
		(*el).Draw(screen)
	}
}

func (ui *UI) DrawDebugUI(screen *ebiten.Image) {
	if !ui.debugUI {
		return
	}

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"GridSpacing": settings.MaxInfluenceRadius,
	}

	world := ebiten.NewImageFromImage(
		screen.SubImage(image.Rect(0, 0, settings.WorldWidth, settings.WorldHeight)),
	)
	op.Images[0] = world

	screen.DrawRectShader(settings.WorldWidth, settings.WorldHeight, gridShader, op)
}
