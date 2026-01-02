package elems

import (
	_ "embed"
	"image/color"
	"life/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ScrollMeter struct {
	position   Position
	dimensions Position

	listener func(float64, float64) float64

	bounds [2]float64
	Value  *float64
	// text *TextElement
}

func NewScrollMeter(x, y, w, h int, listener func(float64, float64) float64, value *float64, high, low float64) *ScrollMeter {
	return &ScrollMeter{
		position:   Pos(x, y),
		dimensions: Pos(w, h),

		bounds:   [2]float64{high, low},
		listener: listener,

		Value: value,
	}
}

func (s *ScrollMeter) Width() int {
	return s.dimensions.X
}

func (s *ScrollMeter) Update(cx, cy int, dw float64, lmb bool) {
	if (cx > s.position.X) && (cx < s.position.X+s.dimensions.X) && (cy > s.position.Y) && (cy < s.position.Y+s.dimensions.Y) {
		*s.Value = utils.Clamp(s.listener(*s.Value, dw), s.bounds[0], s.bounds[1])
	}
}

func (s *ScrollMeter) computeColour() color.Color {
	// Scale between -1 and 1
	centralisedValue := *s.Value - s.bounds[0]
	centralisedValue /= (s.bounds[1] - s.bounds[0]) / 2
	centralisedValue -= 1

	return color.RGBA{
		uint8(utils.Clamp(-centralisedValue, 0, 1) * 200),
		uint8(utils.Clamp(centralisedValue, 0, 1) * 200),
		0,
		0,
	}
}

func (s *ScrollMeter) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(s.position.X), float32(s.position.Y),
		float32(s.dimensions.X), float32(s.dimensions.Y),
		s.computeColour(), true,
	)
}
