package elems

import "github.com/hajimehoshi/ebiten/v2"

type Element interface {
	Update(int, int, float64, bool)
	Draw(*ebiten.Image)
	Width() int
}

type Position struct {
	X, Y int
}

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}
