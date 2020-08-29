package world

import (
	"fmt"
	"github.com/ojrac/opensimplex-go"
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

	n := opensimplex.NewNormalized(0)

	h := SkyLevel - UnderGroundLevel
	for i, block := range c.Data {
		x, y := BlockIndexToPosition(i).Values()
		v := n.Eval2(float64(coords.X*ChunkSize+uint32(x))/100, 0)
		yc := int64(coords.Y)*ChunkSize + int64(y)
		lvl := UnderGroundLevel + int64(v*float64(h))
		if yc < lvl {
			block.Kind = 1
		}
		c.Data[i] = block
	}
	fmt.Printf("generated chunk %d,%d\n", coords.X, coords.Y)
	p.Chunks[coords] = c
}
