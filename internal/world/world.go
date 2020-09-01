package world

import (
	"github.com/hajimehoshi/ebiten"
	"sync"
)

type ColorOffset struct {
	R float32
	G float32
	B float32
}

type Block struct {
	Kind    uint16
	Solid   bool
	Offset  ColorOffset
	TexMain *ebiten.Image
	TexDeco *ebiten.Image
}

type Chunk struct {
	Data [ChunkSize * ChunkSize]Block
}

type Planet struct {
	Size   uint32
	Chunks map[ChunkPosition]*Chunk
	Lock   sync.Mutex
}

func (p *Planet) GetBlock(c Coordinates) *Block {
	p.Lock.Lock()
	ch := p.Chunks[c.ToChunkPosition()]
	p.Lock.Unlock()
	if ch == nil {
		return nil
	}
	return &ch.Data[BlockPositionToIndex(c.ToRelativeBlockPosition())]
}

type Model struct {
	Camera *Camera
	Player *Character
	Planet *Planet
	Debug  bool
}
