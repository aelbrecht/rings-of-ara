package blocks

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/ojrac/opensimplex-go"
	"image"
	"math/rand"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
	"time"
)

var n = opensimplex.NewNormalized(0)

// enqueue chunk generation
func ChunkGenerator(w *world.Model) {
	for {
		time.Sleep(100 * time.Millisecond)
		for _, chPos := range w.Camera.VisibleChunks() {
			GenerateChunk(w.Planet, chPos)
		}
	}
}

func surfaceLevel(c world.Coordinates) int64 {
	h := world.HillsLevel - world.GroundLevel
	v := n.Eval2(float64(c.X)/3000, 0)
	v2 := n.Eval2(float64(c.X)/300, 0)
	v3 := n.Eval2(float64(c.X)/30, 0)
	lvl := world.GroundLevel + int64(v*float64(h)+v2*float64(world.ChunkSize)+v3*float64(world.ChunkSize/2))
	return lvl
}

func TheoreticalSolidType(coords world.Coordinates) bool {
	lvl := surfaceLevel(coords)
	if coords.Y <= lvl {
		return true
	}
	return false
}

func GenerateBlockKind(coords world.Coordinates) (uint16, *ebiten.Image) {
	lvl := surfaceLevel(coords)
	if coords.Y-1 == lvl {
		if rand.Float64() < 0.3 {
			r := rand.Intn(6)
			t := textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{0 + 10*r, 10},
				Max: image.Point{10 + 10*r, 20},
			}).(*ebiten.Image)
			return 3, t
		}
	} else if coords.Y == lvl {
		var t *ebiten.Image
		left := !TheoreticalSolidType(world.Coordinates{coords.X - 1, coords.Y})
		right := !TheoreticalSolidType(world.Coordinates{coords.X + 1, coords.Y})
		if right && left {
			t = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{50, 0},
				Max: image.Point{60, 10},
			}).(*ebiten.Image)
		} else if right {
			t = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{30, 0},
				Max: image.Point{40, 10},
			}).(*ebiten.Image)
		} else if left {
			t = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{40, 0},
				Max: image.Point{50, 10},
			}).(*ebiten.Image)
		} else {
			r := rand.Intn(3)
			t = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{0 + 10*r, 0},
				Max: image.Point{10 + 10*r, 10},
			}).(*ebiten.Image)
		}
		return 2, t
	} else if coords.Y < lvl {
		r := rand.Intn(3)
		t := textures.TileSetUnderground.SubImage(image.Rectangle{
			Min: image.Point{10 + 10*r, 0},
			Max: image.Point{20 + 10*r, 10},
		}).(*ebiten.Image)
		return 1, t
	}
	return 0, nil
}

func GenerateChunk(p *world.Planet, coords world.ChunkPosition) {
	p.Lock.Lock()
	c := p.Chunks[coords]
	p.Lock.Unlock()
	if c != nil {
		return
	}
	c = &world.Chunk{}

	for i, block := range c.Data {
		bl := world.BlockIndexToPosition(i)
		block.Solid = true
		block.Kind, block.Texture = GenerateBlockKind(world.CombineChunkBlockPosition(coords, bl))
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
