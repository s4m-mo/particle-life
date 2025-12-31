package ui

import (
	_ "embed"
	"life/particle"
	"life/settings"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shader.kage
	shaderSource []byte

	shader *ebiten.Shader
)

func init() {
	var err error

	shader, err = ebiten.NewShader(shaderSource)
	if err != nil {
		log.Fatal(err)
	}
}

type UI struct {
	particles *particle.ParticleSet

	wasDown bool
}

func NewUI(particles *particle.ParticleSet) *UI {
	return &UI{
		particles: particles,
	}
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
