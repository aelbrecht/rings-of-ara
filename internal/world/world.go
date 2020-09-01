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

func (p *Planet) GetChunk(c Coordinates) *Chunk {
	p.Lock.Lock()
	ch := p.Chunks[c.ToChunkPosition()]
	p.Lock.Unlock()
	return ch
}

func (p *Planet) GetBlock(c Coordinates) *Block {
	ch := p.GetChunk(c)
	if ch == nil {
		return nil
	}
	return &ch.Data[c.ToRelativeBlockPosition().Index()]
}

type Model struct {
	Camera *Camera
	Player *Character
	Planet *Planet
	Debug  bool
}
