package world

import "rings-of-ara/internal/textures"

type Character struct {
	Step      int
	Settled   bool
	Mask      CharacterMask
	Pos       Coordinates
	Target    Coordinates
	Aiming    Vector
	Vel       Vector
	Direction int
	Tex       *textures.TileSet
}