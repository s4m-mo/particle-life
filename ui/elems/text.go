package elems

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	FontSize       = 18
	FontLineHeight = FontSize + 8
)

var fontface *text.GoTextFaceSource

//go:embed font.ttf
var interFaceSource []byte

func init() {
	f, err := text.NewGoTextFaceSource(bytes.NewReader(interFaceSource))
	if err != nil {
		log.Fatal(err)
	}

	fontface = f
}

type TextElement struct {
	content  string
	position Position

	Size float64
}

func NewTextElement(text string, x, y int) *TextElement {
	return &TextElement{
		content:  text,
		position: Pos(x, y),
		Size:     FontSize,
	}
}

func (t *TextElement) Update(cx, cy int, dw float64, lmb bool) {}

func (t *TextElement) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	geom := ebiten.GeoM{}
	geom.Translate(float64(t.position.X), float64(t.position.Y))

	op.GeoM = geom

	text.Draw(screen, t.content, &text.GoTextFace{Source: fontface, Size: t.Size}, op)
}

func (t *TextElement) Width() int {
	w, _ := text.Measure(t.content+" ", &text.GoTextFace{Source: fontface, Size: FontSize}, FontLineHeight)
	return int(w)
}
