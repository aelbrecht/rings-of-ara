package world

import "github.com/hajimehoshi/ebiten"

type Character struct {
	Step      int
	Settled   bool
	Mask      CharacterMask
	Pos       Coordinates
	Target    Coordinates
	Aiming    Vector
	Vel       Vector
	Draw      func(*Character, *Model, *ebiten.Image)
	Direction int
}
