package elems

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ButtonPaddingY = 3
	ButtonPaddingX = 5
)

type Button struct {
	content    string
	position   Position
	dimensions Position

	colour        color.Color
	hoveredColour color.Color
	pressedColour color.Color

	hovered         bool
	pressed         bool
	wasJustPressed  bool
	pressedListener func()

	text *TextElement
}

func NewButton(s string, x, y, w, h int, listener func(), col, hCol, pCol color.Color) *Button {
	tw, th := text.Measure(s, &text.GoTextFace{Source: fontface, Size: FontSize}, FontLineHeight)

	var dims Position
	var txt *TextElement

	if w == -1 || h == -1 {
		dims = Pos(int(tw)+ButtonPaddingX*2, int(th)+ButtonPaddingY*2)
		txt = NewTextElement(s, x+ButtonPaddingX, y+ButtonPaddingY)
	} else {
		dims = Pos(w, h)
		txt = NewTextElement(
			s,
			int(float64(x)+float64(w/2)-(tw/2)),
			int(float64(y)+float64(h/2)-(th/2)),
		)
	}

	return &Button{
		content:    s,
		position:   Pos(x, y),
		dimensions: dims,

		colour:        col,
		hoveredColour: hCol,
		pressedColour: pCol,

		pressedListener: listener,
		text:            txt,
	}
}

func (b *Button) Width() int {
	return b.dimensions.X
}

func (b *Button) Update(cx, cy int, dw float64, lmb bool) {
	b.hovered = (cx > b.position.X) && (cx < b.position.X+b.dimensions.X) && (cy > b.position.Y) && (cy < b.position.Y+b.dimensions.Y)
	b.pressed = b.hovered && lmb

	if b.pressed {
		if !b.wasJustPressed {
			b.pressedListener()
		}

		b.wasJustPressed = true
	} else {
		b.wasJustPressed = false
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	c := b.colour

	switch {
	case b.pressed:
		c = b.pressedColour
	case b.hovered:
		c = b.hoveredColour
	}

	vector.FillRect(
		screen,
		float32(b.position.X), float32(b.position.Y),
		float32(b.dimensions.X), float32(b.dimensions.Y),
		c, true,
	)

	b.text.Draw(screen)
}
