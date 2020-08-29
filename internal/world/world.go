package world

type Block struct {
	Kind uint16
}

type Chunk struct {
	Data [ChunkSize * ChunkSize]Block
}

type Planet struct {
	Size   uint32
	Chunks map[ChunkPosition]*Chunk
}

func (p *Planet) GetBlock(c Coordinates) *Block {
	ch := p.Chunks[c.ToChunkPosition()]
	if ch == nil {
		return nil
	}
	return &ch.Data[BlockPositionToIndex(c.ToRelativeBlockPosition())]
}

type Model struct {
	Camera *Camera
	Player *Character
	Planet *Planet
}
