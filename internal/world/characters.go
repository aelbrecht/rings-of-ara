package world

import "github.com/hajimehoshi/ebiten"

type Character struct {
	Step      int
	Pos       Coordinates
	Vel       Vector
	Draw      func(*Character, *Model, *ebiten.Image)
	Direction int
}
