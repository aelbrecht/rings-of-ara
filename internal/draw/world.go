package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
	"sync"
)

func init() {
	chunkBuffer = make(map[world.ChunkPosition]*ebiten.Image)
	chunkBufferArray = make([]byte, world.ChunkPointSize*world.ChunkPointSize*4)
}

var chunkBufferLock sync.Mutex
var chunkBuffer map[world.ChunkPosition]*ebiten.Image

var chunkBufferArray []byte

func Chunk(ch *world.Chunk, pos world.ChunkPosition, dst *ebiten.Image, debug bool) {

	for i, _ := range chunkBufferArray {
		chunkBufferArray[i] = 0
	}

	w := world.ChunkPointSize

	for i := range ch.Data {
		rel := world.BlockIndexToPosition(i)
		x, y := rel.Values()
		xp := x * world.BlockSize
		yp := y * world.BlockSize

		if ch.Data[i].Kind == 0 {
			cc := color.RGBA{228, 241, 254, 255}
			for xx := 0; xx < world.BlockSize; xx++ {
				for yy := 0; yy < world.BlockSize; yy++ {
					SetPixel(xp+xx, yp+yy, w, cc, chunkBufferArray)
				}
			}
		} else {
			tex := textures.UndergroundTileMap
			lb := 0
			ub := 1
			row := 0
			switch ch.Data[i].Kind {
			case 1:
				tex = textures.UndergroundTileMap
				lb = 1
				ub = 4
				row = 0
				break
			case 2:
				tex = textures.GrasslandTileMap
				lb = 0
				ub = 3
				row = 0
				co := world.CombineChunkBlockPosition(pos, rel)
				co.X -= 1
				left := world.TheoreticalSolidType(co)
				co.X += 2
				right := world.TheoreticalSolidType(co)
				if !left && !right {
					lb = 5
					ub = 6
				} else if !left {
					lb = 4
					ub = 5
				} else if !right {
					lb = 3
					ub = 4
				}
				break
			case 3:
				tex = textures.GrasslandTileMap
				lb = 0
				ub = 7
				row = 1
				break
			}
			r := lb + rand.Intn(ub-lb)
			tile := tex.Tile(r, row)
			img := tile.Image()
			for xx := 0; xx < world.BlockSize; xx++ {
				for yy := 0; yy < world.BlockSize; yy++ {
					o := GetPixelOffset(xp+xx, yp+(world.BlockSize-yy-1), w)
					o2 := yy*img.Stride + xx*4
					chunkBufferArray[o] = img.Pix[o2]
					chunkBufferArray[o+1] = img.Pix[o2+1]
					chunkBufferArray[o+2] = img.Pix[o2+2]
					chunkBufferArray[o+3] = img.Pix[o2+3]
				}
			}
		}

		if debug {
			c := color.RGBA{100, 100, 100, 255}
			for i := xp; i < xp+world.BlockSize; i++ {
				SetPixel(i*world.BlockSize, yp, w, c, chunkBufferArray)
				SetPixel(i, yp+world.BlockSize-1, w, c, chunkBufferArray)
			}
			for i := yp; i < yp+world.BlockSize; i++ {
				SetPixel(xp, i, w, c, chunkBufferArray)
				SetPixel(xp+world.BlockSize-1, i, w, c, chunkBufferArray)
			}
		}
	}

	if debug {
		Rectangle(world.ChunkPointSize, color.RGBA{255, 0, 0, 50}, chunkBufferArray)
	}

	_ = dst.ReplacePixels(chunkBufferArray)
}

func Planet(w *world.Model, screen *ebiten.Image) {
	for _, coords := range w.Camera.VisibleChunks() {
		chunkBufferLock.Lock()
		if chunkBuffer[coords] == nil {
			chunkBufferLock.Unlock()
			enqueueChunk(coords)
			continue
		}
		chunkTexture := chunkBuffer[coords]
		chunkBufferLock.Unlock()
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(3, -3)
		x, y := w.Camera.ToScreen(coords.ToCoordinates()).ValuesFloat()
		op.GeoM.Translate(x, y)
		screen.DrawImage(chunkTexture, &op)
	}
}
