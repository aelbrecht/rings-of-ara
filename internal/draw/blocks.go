package draw

import (
	"github.com/hajimehoshi/ebiten"
	"rings-of-ara/internal/world"
)

func BlockLayer(w *world.Model, screen *ebiten.Image) {

	// draw visible chunks on screen
	for _, coords := range w.Camera.VisibleChunks() {

		chunk := w.Planet.Chunks[coords]
		if chunk == nil {
			continue
		}

		for i, block := range chunk.Data {

			if block.TexMain == nil {
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

			sub := block.TexMain
			top := block.TexDeco

			screen.DrawImage(sub, &op)

			if top != nil {
				screen.DrawImage(top, &op)
			}
		}
	}
}
