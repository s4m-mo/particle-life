package ui

import (
	_ "embed"
	"image"
	"life/particle"
	"life/settings"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shader.kage
	shaderSource []byte

	//go:embed gridshader.kage
	gridShaderSource []byte

	shader     *ebiten.Shader
	gridShader *ebiten.Shader
)

func init() {
	var err error

	if shader, err = ebiten.NewShader(shaderSource); err != nil {
		log.Fatal(err)
	}
	if gridShader, err = ebiten.NewShader(gridShaderSource); err != nil {
		log.Fatal(err)
	}
}

type UI struct {
	particles *particle.ParticleSet

	wasDown bool
	debugUI bool
}

func NewUI(particles *particle.ParticleSet) *UI {
	return &UI{
		particles: particles,
	}
}

func (ui *UI) ToggleDebugUI() {
	ui.debugUI = !ui.debugUI
}

func (ui *UI) Update() {

}

func (ui *UI) Draw(screen *ebiten.Image) {
	image := ebiten.NewImage(settings.WorldWidth, settings.WorldHeight)

	w, h := image.Bounds().Dx(), image.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	geom := ebiten.GeoM{}
	geom.Translate(settings.ScreenWidth-settings.UIWidth, 0)

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"Cursor": []float32{float32(cx), float32(cy)},
	}

	op.Images[0] = image
	op.GeoM = geom

	screen.DrawRectShader(w, h, shader, op)
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
