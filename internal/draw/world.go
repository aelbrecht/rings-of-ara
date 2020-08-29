package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"rings-of-ara/internal/world"
	"sync"
)

func init() {
	chunkBuffer = make(map[world.ChunkPosition]*ebiten.Image)
	chunkBufferArray = make([]byte, world.ChunkPixelSize*world.ChunkPixelSize*4)
}

var chunkBufferLock sync.Mutex
var chunkBuffer map[world.ChunkPosition]*ebiten.Image

var chunkBufferArray []byte

func Chunk(ch *world.Chunk, dst *ebiten.Image, debug bool) {

	for i, _ := range chunkBufferArray {
		chunkBufferArray[i] = 0
	}

	w := world.ChunkPixelSize
	for i := range ch.Data {

		x, y := world.BlockIndexToPosition(i).Values()
		xp := x * world.BlockPixelSize
		yp := y * world.BlockPixelSize

		cc := color.RGBA{0, 0, 0, 0}
		switch ch.Data[i].Kind {
		case 0:
			cc = color.RGBA{228, 241, 254, 255}
		case 1:
			cc = color.RGBA{0, 200, 0, 255}
			break
		}
		for xx := 0; xx < world.BlockPixelSize; xx++ {
			for yy := 0; yy < world.BlockPixelSize; yy++ {
				SetPixel(xp+xx, yp+yy, w, cc, chunkBufferArray)
			}
		}

		if debug {
			c := color.RGBA{100, 100, 100, 255}
			for i := xp; i < xp+world.BlockPixelSize; i++ {
				SetPixel(i*world.BlockPixelSize, yp, w, c, chunkBufferArray)
				SetPixel(i, yp+world.BlockPixelSize-1, w, c, chunkBufferArray)
			}
			for i := yp; i < yp+world.BlockPixelSize; i++ {
				SetPixel(xp, i, w, c, chunkBufferArray)
				SetPixel(xp+world.BlockPixelSize-1, i, w, c, chunkBufferArray)
			}
		}
	}

	if debug {
		Rectangle(world.ChunkPixelSize, color.RGBA{255, 0, 0, 50}, chunkBufferArray)
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
		op.GeoM.Scale(1, -1)
		x, y := w.Camera.ToScreen(coords.ToCoordinates()).ValuesFloat()
		op.GeoM.Translate(x, y)
		screen.DrawImage(chunkTexture, &op)
	}
}
