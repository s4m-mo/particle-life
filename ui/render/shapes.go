package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Shape interface {
	/* DrawInternal draws the shape to the image in a way that the shader will
	   understand. The values of a colour represent the following:

	   A: radius of a circle*/
	DrawInternal(*ebiten.Image)
}

type Circle struct {
	x, y int
	r    uint8
}

func (c *Circle) DrawInternal(img *ebiten.Image) {
	img.Set(c.x, c.y, color.RGBA{0, 0, 0, c.r})
}
