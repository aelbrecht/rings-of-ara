package draw

import (
	"github.com/hajimehoshi/ebiten"
	"rings-of-ara/internal/world"
)

func BlockLayer(w *world.Model, front *ebiten.Image, back *ebiten.Image) {

	// draw visible chunks on screen
	for _, coords := range w.Camera.VisibleChunks() {

		chunk := w.Planet.Chunks[coords]
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

			if block.TexMain != nil {
				back.DrawImage(block.TexMain, &op)
			}

			if block.TexDeco != nil {
				if block.Offset.R != 0 || block.Offset.G != 0 || block.Offset.B != 0 {
					r := float64(block.Offset.R)
					g := float64(block.Offset.G)
					b := float64(block.Offset.B)
					op.ColorM.Scale(1+r, 1+g, 1+b, 1)
				}
				front.DrawImage(block.TexDeco, &op)
			}
		}
	}
}
