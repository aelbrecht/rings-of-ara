package world

import "github.com/hajimehoshi/ebiten"

type Block struct {
	Kind uint16
}

var blockBuffer, _ = ebiten.NewImage(32*3, 32*3, ebiten.FilterDefault)

type Chunk struct {
	Data [32 * 32]Block
}

type Planet struct {
	Size   uint32
	Chunks map[ChunkPosition]*Chunk
}

type Model struct {
	Camera *Camera
	Player *Character
	Planet *Planet
}
