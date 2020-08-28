package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"rings-of-ara/internal/world"
)

func init() {
	chunkBuffer = make(map[world.ChunkPosition]*ebiten.Image)
	chunkBufferArray = make([]byte, world.ChunkPixelSize*world.ChunkPixelSize*4)
}

var chunkBuffer map[world.ChunkPosition]*ebiten.Image

var chunkBufferArray []byte

func Chunk(ch *world.Chunk, dst *ebiten.Image) {

	for i, _ := range chunkBufferArray {
		chunkBufferArray[i] = 0
	}

	w := world.ChunkPixelSize
	c := color.RGBA{0, 0, 0, 255}
	for i := range ch.Data {
		x, y := world.BlockIndexToPosition(i).Values()
		xp := x*world.BlockPixelSize
		yp := y*world.BlockPixelSize
		for i := xp; i < xp+world.BlockPixelSize; i++ {
			SetPixel(i*world.BlockPixelSize, yp, w, c, chunkBufferArray)
			SetPixel(i, yp+world.BlockPixelSize-1, w, c, chunkBufferArray)
		}
		for i := yp; i < yp+world.BlockPixelSize; i++ {
			SetPixel(xp, i, w, c, chunkBufferArray)
			SetPixel(xp+world.BlockPixelSize-1, i, w, c, chunkBufferArray)
		}
	}

	Rectangle(world.ChunkPixelSize, color.RGBA{255, 0, 0, 255}, chunkBufferArray)

	dst.ReplacePixels(chunkBufferArray)
}

func Planet(w *world.Model, screen *ebiten.Image) {

	for _, coords := range w.Camera.VisibleChunks() {
		if chunkBuffer[coords] == nil {
			enqueueChunk(coords)
			continue
		}
		chunkTexture := chunkBuffer[coords]
		op := ebiten.DrawImageOptions{}
		x, y := coords.ToCoordinates().Sub(w.Camera.Offset()).ValuesFloat()
		op.GeoM.Translate(x, y)
		screen.DrawImage(chunkTexture, &op)
	}
}
