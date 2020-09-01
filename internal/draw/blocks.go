package draw

import (
	"github.com/hajimehoshi/ebiten"
	"rings-of-ara/internal/world"
)

var BlockRenders = 0

// render back walls, solid blocks...
func renderBack(block *world.Block, back *ebiten.Image, op *ebiten.DrawImageOptions) {
	if block.TexMain != nil {
		back.DrawImage(block.TexMain, op)
		BlockRenders++
	}
}

// render foliage, grass...
func renderFront(block *world.Block, front *ebiten.Image, op *ebiten.DrawImageOptions) {
	if block.TexDeco != nil {
		if block.Offset.R != 0 || block.Offset.G != 0 || block.Offset.B != 0 {
			r := float64(block.Offset.R)
			g := float64(block.Offset.G)
			b := float64(block.Offset.B)
			op.ColorM.Scale(1+r, 1+g, 1+b, 1)
		}
		front.DrawImage(block.TexDeco, op)
		BlockRenders++
	}
}

// render a screen of blocks
func renderLayer(w *world.Model, layer *ebiten.Image, draw func(*world.Block, *ebiten.Image, *ebiten.DrawImageOptions)) {

	// draw visible chunks on screen
	for _, coords := range w.Camera.VisibleChunks() {

		w.Planet.Lock.Lock()
		chunk := w.Planet.Chunks[coords]
		w.Planet.Lock.Unlock()
		if chunk == nil {
			continue
		}

		for i, block := range chunk.Data {

			if block.TexMain == nil && block.TexDeco == nil {
				continue
			}

			op := ebiten.DrawImageOptions{}

			// adjust for flipped y axis
			op.GeoM.Translate(0, -world.BlockSize)

			// scale blocks 3x
			op.GeoM.Scale(3, 3)

			// get screen position of block
			pos := world.CombineChunkBlockPosition2(coords, world.BlockIndexToPosition(i))
			x, y := w.Camera.ToScreen(pos).ValuesFloat()
			op.GeoM.Translate(x, y)

			draw(&block, layer, &op)
		}
	}
}

// renders front and back block layers
// front should be drawn in front of characters
// render splits up into two different images to improve performance as successive draw calls are more efficient
func BlockLayer(w *world.Model, front *ebiten.Image, back *ebiten.Image) {
	BlockRenders = 0
	renderLayer(w, front, renderFront)
	renderLayer(w, back, renderBack)
}
