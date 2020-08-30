package world

import (
	"fmt"
	"github.com/ojrac/opensimplex-go"
	"math/rand"
	"time"
)

var n = opensimplex.NewNormalized(0)

// enqueue chunk generation
func ChunkGenerator(w *Model) {
	for {
		time.Sleep(100 * time.Millisecond)
		for _, chPos := range w.Camera.VisibleChunks() {
			w.Planet.GenerateChunk(chPos)
		}
	}
}

func GetBlockType(x int, y int, coords ChunkPosition, m *Model) {
}

func TheoreticalSolidType(coords Coordinates) bool {
	h := HillsLevel - GroundLevel
	v := n.Eval2(float64(coords.X)/3000, 0)
	v2 := n.Eval2(float64(coords.X)/300, 0)
	v3 := n.Eval2(float64(coords.X)/30, 0)
	lvl := GroundLevel + int64(v*float64(h)+v2*float64(ChunkSize)+v3*float64(ChunkSize/2))
	if coords.Y <= lvl {
		return true
	}
	return false
}

func TheoreticalBlockType(coords Coordinates) uint16 {
	h := HillsLevel - GroundLevel
	v := n.Eval2(float64(coords.X)/3000, 0)
	v2 := n.Eval2(float64(coords.X)/300, 0)
	v3 := n.Eval2(float64(coords.X)/30, 0)
	lvl := GroundLevel + int64(v*float64(h)+v2*float64(ChunkSize)+v3*float64(ChunkSize/2))
	if coords.Y == lvl+1 {
		if rand.Float64() < 0.3 {
			return 3
		}
	} else if coords.Y == lvl {
		return 2
	} else if coords.Y < lvl {
		return 1
	}
	return 0
}

func (p *Planet) GenerateChunk(coords ChunkPosition) {
	p.Lock.Lock()
	c := p.Chunks[coords]
	p.Lock.Unlock()
	if c != nil {
		return
	}
	c = &Chunk{}

	for i, block := range c.Data {
		bl := BlockIndexToPosition(i)
		block.Solid = true
		block.Kind = TheoreticalBlockType(CombineChunkBlockPosition(coords, bl))
		if block.Kind == 3 || block.Kind == 0 {
			block.Solid = false
		}
		c.Data[i] = block
	}
	fmt.Printf("generated chunk %d,%d\n", coords.X, coords.Y)

	p.Lock.Lock()
	p.Chunks[coords] = c
	p.Lock.Unlock()
}
