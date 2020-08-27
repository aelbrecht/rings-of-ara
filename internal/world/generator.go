package world

func (p *Planet) GenerateChunk(coords ChunkPosition) {
	c := p.Chunks[coords]
	if c != nil {
		return
	}
	c = &Chunk{}
	p.Chunks[coords] = c
}
