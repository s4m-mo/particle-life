package particle

import (
	_ "embed"
	"image/color"
	"life/settings"
	"life/utils"
	"log"
	"math"
	"math/rand"

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

func computeRandomAttractionMatrix(variants int) [][]float64 {
	out := make([][]float64, variants)
	for i := range out {
		out[i] = make([]float64, variants)

		for j := range out[i] {
			out[i][j] = rand.Float64()*2 - 1
			//out[i][j] = utils.Clamp(rand.NormFloat64()/3, -1, 1)
		}
	}

	return out
}

type DividedSpace [][][]*Particle

type ParticleSet struct {
	particles []*Particle

	nVariants        int
	variantColours   []color.Color
	AttractionMatrix [][]float64

	nWidth  int
	nHeight int
	space   DividedSpace
}

func NewRandomParticleSet(n int, variants int) *ParticleSet {
	ps := &ParticleSet{
		particles:        make([]*Particle, n),
		AttractionMatrix: computeRandomAttractionMatrix(variants),
		nVariants:        variants,
	}

	ps.RegenerateUniformPoints(n, variants)

	ps.variantColours = computeColours(variants)
	ps.computeSpaceDivisions()

	return ps
}

func NewCentredParticleSet(n int, variants int) *ParticleSet {
	ps := &ParticleSet{
		particles:        make([]*Particle, n),
		AttractionMatrix: computeRandomAttractionMatrix(variants),
		nVariants:        variants,
	}

	ps.RegenerateCentralPoints(n, variants)

	ps.variantColours = computeColours(variants)
	ps.computeSpaceDivisions()

	return ps
}

func (ps *ParticleSet) RegenerateUniformPoints(n int, variants int) {
	for i := 0; i < n; i++ {
		ps.particles[i] = &Particle{
			x:       float64(rand.Intn(settings.WorldWidth)),
			y:       float64(rand.Intn(settings.WorldHeight)),
			vx:      rand.NormFloat64(),
			vy:      rand.NormFloat64(),
			variant: uint(rand.Intn(variants)),
		}
	}
}

func (ps *ParticleSet) RegenerateCentralPoints(n int, variants int) {
	for i := 0; i < n; i++ {
		ps.particles[i] = &Particle{
			x:       math.Mod(centralRandPoint(settings.WorldWidth), settings.WorldWidth),
			y:       math.Mod(centralRandPoint(settings.WorldHeight), settings.WorldHeight),
			vx:      rand.NormFloat64(),
			vy:      rand.NormFloat64(),
			variant: uint(rand.Intn(variants)),
		}
	}
}

func (ps *ParticleSet) RegenerateOuterPoints(n int, variants int) {
	for i := 0; i < n; i++ {
		ps.particles[i] = &Particle{
			x:       math.Mod(centralRandPoint(settings.WorldWidth)+(settings.WorldWidth/2), settings.WorldWidth),
			y:       math.Mod(centralRandPoint(settings.WorldHeight)+(settings.WorldHeight/2), settings.WorldHeight),
			vx:      rand.NormFloat64(),
			vy:      rand.NormFloat64(),
			variant: uint(rand.Intn(variants)),
		}
	}
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

func (ps *ParticleSet) SetAttractionMatrix(v float64) {
	for i := range ps.AttractionMatrix {
		for j := range ps.AttractionMatrix[i] {
			ps.AttractionMatrix[i][j] = v
		}
	}
}

func (ps *ParticleSet) RegenerateAttractionMatrix() {
	for i := range ps.AttractionMatrix {
		for j := range ps.AttractionMatrix[i] {
			ps.AttractionMatrix[i][j] = rand.Float64()*2 - 1
			//ps.AttractionMatrix[i][j] = utils.Clamp(rand.NormFloat64()/3, -1, 1)
		}
	}
}

func generateSpaceDivisionMatrix(nWidth, nHeight int) DividedSpace {

	space := make([][][]*Particle, nWidth)
	for i := range space {
		space[i] = make([][]*Particle, nHeight)

		for j := range space[i] {
			space[i][j] = make([]*Particle, 0)
		}
	}

	return space
}

func (ps *ParticleSet) computeSpaceDivisions() {
	ps.nWidth = int(math.Ceil(settings.WorldWidth / settings.MaxInfluenceRadius))
	ps.nHeight = int(math.Ceil(settings.WorldHeight / settings.MaxInfluenceRadius))

	sp := generateSpaceDivisionMatrix(ps.nWidth, ps.nHeight)
	ps.space = sp

	ps.computeSpaceDivisionLocations()
}

func (ps *ParticleSet) computeSpaceDivisionLocations() {
	for _, v := range ps.particles {
		nx, ny := v.CurrentQuadrant()
		ps.space[nx][ny] = append(ps.space[nx][ny], v)
	}
}

// Updates particle based on surrounding particles and returns the particle's new quadrant
func (ps *ParticleSet) computeParticleUpdate(o *Particle, dt float64) (int, int) {
	// Compute which quadrants are needed to be searched
	quadrants := o.FindInfluencingQuadrants()
	home := quadrants[0]

	forceX := 0.0
	forceY := 0.0

	// Calculate influcences
	for _, quadrant := range quadrants {
		// Fix negative quadrant indices and overflows on the other side
		nx := quadrant[0]
		ny := quadrant[1]

		for nx < 0 {
			nx += ps.nWidth
		}

		for ny < 0 {
			ny += ps.nHeight
		}

		nx = nx % ps.nWidth
		ny = ny % ps.nHeight

		items := ps.space[nx][ny]

		// Handle wrap-around coordinates
		var wrappingOffset [2]float64

		if int(math.Abs(float64(nx-home[0]))) == ps.nWidth-1 {
			wrappingOffset[0] = settings.MaxInfluenceRadius * float64(ps.nWidth)
			if home[0] < nx {
				wrappingOffset[0] *= -1
			}
		}

		if int(math.Abs(float64(ny-home[1]))) == ps.nHeight-1 {
			wrappingOffset[1] = settings.MaxInfluenceRadius * float64(ps.nHeight)
			if home[1] < ny {
				wrappingOffset[1] *= -1
			}
		}

		for _, i := range items {
			dx := i.x - o.x + wrappingOffset[0]
			dy := i.y - o.y + wrappingOffset[1]

			// dist := (dx + dy) // Manhattan Distance
			dist := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

			f := PolyForce(dist, ps.AttractionMatrix[o.variant][i.variant]) * settings.UniversalForceMultiplier

			// Add components of force in each direction
			if dist != 0 {
				forceX += f * (dx / dist)
				forceY += f * (dy / dist)
			}
		}
	}

	// Apply forces to velocity
	o.vx += forceX * dt
	o.vy += forceY * dt

	// Apply velocity
	o.Update()

	// Compute new quadrant
	return o.CurrentQuadrant()
}

func (ps *ParticleSet) Update(dt float64) {
	cx, cy := ebiten.CursorPosition()

	isCursorPressedRepel := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && (cx < settings.WorldWidth)
	isCursorPressedAttract := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && (cx < settings.WorldWidth)

	newQuads := make(chan QuadInfo, len(ps.particles))

	for _, p := range ps.particles {
		go func(outChan chan<- QuadInfo, delta, cursorX, cursorY float64, repel, attract bool) {

			// Repel from cursor
			if SquaredEuclideanDistance(p.x, p.y, cursorX, cursorY) < settings.CursorForceRadiusSquared {
				switch {
				case attract:
					p.vx += (cursorX - p.x) * settings.CursorRepelForce
					p.vy += (cursorY - p.y) * settings.CursorRepelForce
				case repel:
					p.vx += (p.x - cursorX) * settings.CursorRepelForce
					p.vy += (p.y - cursorY) * settings.CursorRepelForce
				}
			}

			nx, ny := ps.computeParticleUpdate(p, delta)
			outChan <- QuadInfo{X: nx, Y: ny, P: p}
		}(newQuads, dt, float64(cx), float64(cy), isCursorPressedRepel, isCursorPressedAttract)
	}

	// Generate new space
	newSpace := generateSpaceDivisionMatrix(ps.nWidth, ps.nHeight)

	// Receive new quadrants from the goroutines
	for range len(ps.particles) {
		msg := <-newQuads
		newSpace[msg.X][msg.Y] = append(newSpace[msg.X][msg.Y], msg.P)
	}

	// Replace old space
	ps.space = newSpace
}

func (ps *ParticleSet) Draw(screen *ebiten.Image) {
	image := ebiten.NewImage(settings.WorldWidth, settings.WorldHeight)

	for _, v := range ps.particles {
		(*v).Draw(image, ps.variantColours[(*v).variant])
	}

	w, h := image.Bounds().Dx(), image.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"C":   []float64{float64(cx), float64(cy)},
		"Cr2": settings.CursorForceRadiusSquared,
	}

	op.Images[0] = image

	screen.DrawRectShader(w, h, shader, op)
}
