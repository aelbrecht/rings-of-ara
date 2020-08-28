package world

import (
	"fmt"
	"time"
)

// enqueue chunk generation
func ChunkGenerator(w *Model) {
	for {
		time.Sleep(100 * time.Millisecond)
		for _, chPos := range w.Camera.VisibleChunks() {
			w.Planet.GenerateChunk(chPos)
		}
	}
}

func (p *Planet) GenerateChunk(coords ChunkPosition) {
	c := p.Chunks[coords]
	if c != nil {
		return
	}
	c = &Chunk{}
	for i, block := range c.Data {
		block.Kind = 1
		c.Data[i] = block
	}
	fmt.Printf("generated chunk %d,%d\n", coords.X, coords.Y)
	p.Chunks[coords] = c
}
