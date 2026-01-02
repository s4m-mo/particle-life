package elems

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Box struct {
	position   Position
	dimensions Position

	color color.Color
	// text *TextElement
}

func NewBox(x, y, w, h int, col color.Color) *Box {
	return &Box{
		position:   Pos(x, y),
		dimensions: Pos(w, h),
		color:      col,
	}
}

func (s *Box) Width() int {
	return s.dimensions.X
}

func (s *Box) Update(cx, cy int, dw float64, lmb bool) {}

func (s *Box) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		float32(s.position.X), float32(s.position.Y),
		float32(s.dimensions.X), float32(s.dimensions.Y),
		s.color, true,
	)
}
