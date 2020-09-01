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

func GenerateBlockKind(coords world.Coordinates) (uint16, *ebiten.Image, *ebiten.Image, world.ColorOffset) {
	offset := world.ColorOffset{}
	lvl := surfaceLevel(coords)
	if coords.Y-1 == lvl {
		left := !TheoreticalSolidType(world.Coordinates{coords.X - 1, coords.Y - 1})
		right := !TheoreticalSolidType(world.Coordinates{coords.X + 1, coords.Y - 1})
		v := n.Eval2(float64(coords.X)/50, 0)
		v1 := n.Eval2(float64(coords.X)/200, 0)
		offset.G = float32(0.25*v + 0.75)
		offset.R = float32(0.5 * v1)
		r1 := rand.Intn(2)
		r2 := rand.Intn(6)
		if rand.Float32() < 0.6 {
			r2 = 7
		}
		var t1 *ebiten.Image
		if left && right {
			t1 = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{20, 10},
				Max: image.Point{30, 20},
			}).(*ebiten.Image)
		} else if left {
			t1 = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{40, 10},
				Max: image.Point{50, 20},
			}).(*ebiten.Image)
		} else if right {
			t1 = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{30, 10},
				Max: image.Point{40, 20},
			}).(*ebiten.Image)
		} else {
			t1 = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{0 + 10*r1, 10},
				Max: image.Point{10 + 10*r1, 20},
			}).(*ebiten.Image)
		}
		t2 := textures.TileSetGrassland.SubImage(image.Rectangle{
			Min: image.Point{0 + 10*r2, 20},
			Max: image.Point{10 + 10*r2, 30},
		}).(*ebiten.Image)
		return 3, t2, t1, offset
	} else if coords.Y == lvl {
		v := n.Eval2(float64(coords.X)/50, 0)
		v1 := n.Eval2(float64(coords.X)/200, 0)
		offset.G = float32(0.25*v + 0.75)
		offset.R = float32(0.5 * v1)
		var front *ebiten.Image
		var back *ebiten.Image
		left := !TheoreticalSolidType(world.Coordinates{coords.X - 1, coords.Y})
		right := !TheoreticalSolidType(world.Coordinates{coords.X + 1, coords.Y})
		if right && left {
			front = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{50, 0},
				Max: image.Point{60, 10},
			}).(*ebiten.Image)
			back = textures.TileSetUnderground.SubImage(image.Rectangle{
				Min: image.Point{60, 0},
				Max: image.Point{70, 10},
			}).(*ebiten.Image)
		} else if right {
			front = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{30, 0},
				Max: image.Point{40, 10},
			}).(*ebiten.Image)
			back = textures.TileSetUnderground.SubImage(image.Rectangle{
				Min: image.Point{40, 0},
				Max: image.Point{50, 10},
			}).(*ebiten.Image)
		} else if left {
			front = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{40, 0},
				Max: image.Point{50, 10},
			}).(*ebiten.Image)
			back = textures.TileSetUnderground.SubImage(image.Rectangle{
				Min: image.Point{50, 0},
				Max: image.Point{60, 10},
			}).(*ebiten.Image)
		} else {
			r1 := rand.Intn(3)
			r2 := rand.Intn(3)
			front = textures.TileSetGrassland.SubImage(image.Rectangle{
				Min: image.Point{0 + 10*r1, 0},
				Max: image.Point{10 + 10*r1, 10},
			}).(*ebiten.Image)
			back = textures.TileSetUnderground.SubImage(image.Rectangle{
				Min: image.Point{10 + 10*r2, 0},
				Max: image.Point{20 + 10*r2, 10},
			}).(*ebiten.Image)
		}
		return 2, back, front, offset
	} else if coords.Y < lvl {
		r := rand.Intn(3)
		t := textures.TileSetUnderground.SubImage(image.Rectangle{
			Min: image.Point{10 + 10*r, 0},
			Max: image.Point{20 + 10*r, 10},
		}).(*ebiten.Image)
		return 1, t, nil, offset
	}
	return 0, nil, nil, offset
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
		block.Kind, block.TexMain, block.TexDeco, block.Offset = GenerateBlockKind(world.CombineChunkBlockPosition(coords, bl))
		if block.Kind == 3 || block.Kind == 0 {
			block.Solid = false
		}
		c.Data[i] = block
	}

	p.Lock.Lock()
	p.Chunks[coords] = c
	p.Lock.Unlock()
}
