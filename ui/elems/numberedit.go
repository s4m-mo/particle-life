package elems

import (
	_ "embed"
	"fmt"
	"image/color"
	"life/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type NumberEdit struct {
	position   Position
	dimensions Position

	listener func(float64, float64) float64

	bounds [2]float64
	Value  *float64
	text   *TextElement
	col    color.Color
}

func NewNumberEdit(x, y, w, h int, listener func(float64, float64) float64, value *float64, high, low float64, col color.Color) *NumberEdit {
	return &NumberEdit{
		position:   Pos(x, y),
		dimensions: Pos(w, h),

		bounds:   [2]float64{high, low},
		listener: listener,

		Value: value,
		col:   col,
	}
}

func (n *NumberEdit) Width() int {
	return n.dimensions.X
}

func (n *NumberEdit) Update(cx, cy int, dw float64, lmb bool) {
	if (cx > n.position.X) && (cx < n.position.X+n.dimensions.X) && (cy > n.position.Y) && (cy < n.position.Y+n.dimensions.Y) {
		*n.Value = utils.Clamp(n.listener(*n.Value, dw), n.bounds[0], n.bounds[1])
	}
}

func (n *NumberEdit) Draw(screen *ebiten.Image) {
	txt := fmt.Sprint(n.Value)
	tw, _ := text.Measure(txt, &text.GoTextFace{Source: fontface, Size: FontSize}, FontLineSpacing)

	t := NewTextElement(txt, n.position.X+int(float64(n.dimensions.X)-tw)/2, n.position.Y)
	t.Draw(screen)

	vector.FillRect(
		screen,
		float32(n.position.X), float32(n.position.Y),
		float32(n.dimensions.X), float32(n.dimensions.Y),
		n.col, true,
	)
}
