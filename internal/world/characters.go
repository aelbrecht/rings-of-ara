package world

import (
	"github.com/hajimehoshi/ebiten"
)

type Character struct {
	Step       int
	Settled    bool
	Mask       CharacterMask
	Pos        Coordinates
	Target     Coordinates
	Aiming     Vector
	Vel        Vector
	Direction  int
	Appearance *CharacterAppearance
}

type CharacterFeature struct {
	Darken         float64
	Texture        *ebiten.Image
	Rotation       float64
	OffsetX        float64
	OffsetY        float64
	AnchorX        float64
	AnchorY        float64
	ChildrenBefore []CharacterFeature
	ChildrenBehind []CharacterFeature
}

type CharacterAppearance struct {
	Body CharacterFeature
}
