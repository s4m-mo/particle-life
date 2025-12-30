package particle

import (
	_ "embed"
	"image/color"
	"life/settings"
	"life/utils"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/shader.kage
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

func centralRandPoint(max float64) float64 {
	val := rand.NormFloat64()*float64(max)/8 + max/2
	return utils.Clamp(val, 0, max-1)
}

func computeColours(variants int) []color.Color {
	variantColours := make([]color.Color, variants)
	for i := 0; i < variants; i++ {
		variantColours[i] = utils.HSVToRGB(
			float64(i)*360/float64(variants),
			settings.ColourSaturation,
			settings.ColourValue,
		)
	}

	return variantColours
}

type ParticleSet struct {
	particles []*Particle

	variantColours []color.Color
}

func NewRandomParticleSet(n int, variants int) *ParticleSet {
	ps := &ParticleSet{
		particles: make([]*Particle, n),
	}

	for i := 0; i < n; i++ {
		ps.particles[i] = &Particle{
			x:       float64(rand.Intn(settings.WorldWidth)),
			y:       float64(rand.Intn(settings.WorldHeight)),
			vx:      rand.NormFloat64(),
			vy:      rand.NormFloat64(),
			variant: uint(rand.Intn(variants)),
		}
	}

	ps.variantColours = computeColours(variants)

	return ps
}

func NewCentredParticleSet(n int, variants int) *ParticleSet {
	ps := &ParticleSet{
		particles: make([]*Particle, n),
	}

	for i := 0; i < n; i++ {
		ps.particles[i] = &Particle{
			x:       centralRandPoint(settings.WorldWidth),
			y:       centralRandPoint(settings.WorldHeight),
			vx:      rand.NormFloat64(),
			vy:      rand.NormFloat64(),
			variant: uint(rand.Intn(variants)),
		}
	}

	ps.variantColours = computeColours(variants)

	return ps
}

func (ps *ParticleSet) SetNVariants(n int) {
	ps.variantColours = computeColours(n)
}

func (ps *ParticleSet) GetVariantColours() []color.Color {
	return ps.variantColours
}

func (ps *ParticleSet) GetParticles() []*Particle {
	return ps.particles
}

func (ps *ParticleSet) Update() {
	for _, v := range ps.particles {
		(*v).Update()
	}
}

func (ps *ParticleSet) Draw(screen *ebiten.Image) {
	image := ebiten.NewImageFromImage(screen)

	for _, v := range ps.particles {
		(*v).Draw(image, ps.variantColours[(*v).variant])
	}

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"Cursor": []float32{float32(cx), float32(cy)},
	}

	op.Images[0] = image

	screen.DrawRectShader(w, h, shader, op)
}
