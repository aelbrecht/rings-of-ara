package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"rings-of-ara/internal/world"
)

var characterBuffer *ebiten.Image

func init() {
	characterBuffer, _ = ebiten.NewImage(128, 128, ebiten.FilterDefault)
}

func drawFeature(feature *world.CharacterFeature, rx float64, ry float64, rr float64) {

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-feature.AnchorX, -feature.AnchorY)
	op.GeoM.Rotate(feature.Rotation)
	op.GeoM.Translate(feature.OffsetX, feature.OffsetY)
	op.GeoM.Rotate(rr)
	op.GeoM.Translate(rx, ry)
	if feature.Darken != 0 {
		s := 1 - feature.Darken
		op.ColorM.Scale(s, s, s, 1)
	}
	if feature.ChildrenBehind != nil {
		for i := 0; i < len(feature.ChildrenBehind); i++ {
			drawFeature(&feature.ChildrenBehind[i], rx+feature.OffsetX, ry+feature.OffsetY, rr+feature.Rotation)
		}
	}
	if feature.Texture != nil {
		characterBuffer.DrawImage(feature.Texture, &op)
	}
	if feature.ChildrenBefore != nil {
		for i := 0; i < len(feature.ChildrenBefore); i++ {
			drawFeature(&feature.ChildrenBefore[i], rx+feature.OffsetX, ry+feature.OffsetY, rr+feature.Rotation)
		}
	}
}

func DrawCharacter(ch *world.Character, world *world.Model, screen *ebiten.Image) {

	characterBuffer.Fill(color.RGBA{10, 10, 10, 25})

	drawFeature(&ch.Appearance.Body, 0, 0, 0)

	xf, yf := world.Camera.ToScreen(ch.Pos).ValuesFloat()

	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(xf, yf-512)

	screen.DrawImage(characterBuffer, &op)

}
