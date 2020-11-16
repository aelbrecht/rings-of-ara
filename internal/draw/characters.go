package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"rings-of-ara/internal/world"
)

var characterBuffer *ebiten.Image

func init() {
	characterBuffer, _ = ebiten.NewImage(186, 186, ebiten.FilterDefault)
}

func drawFeature(feature *world.CharacterFeature, opGlobal ebiten.GeoM) {

	// Do rotation of self
	rotation := ebiten.GeoM{}
	rotation.Translate(-feature.AnchorX, -feature.AnchorY)
	rotation.Rotate(feature.Rotation)
	translation := ebiten.GeoM{}
	translation.Translate(feature.OffsetX, feature.OffsetY)

	t := ebiten.GeoM{}
	t.Concat(rotation)
	t.Concat(translation)
	t.Concat(opGlobal)

	tt := ebiten.GeoM{}
	tt.Translate(feature.AnchorX, feature.AnchorY)

	tt.Concat(t)

	op := ebiten.DrawImageOptions{}
	op.GeoM = t

	if feature.Darken != 0 {
		s := 1 - feature.Darken
		op.ColorM.Scale(s, s, s, 1)
	}
	if feature.ChildrenBehind != nil {
		for i := 0; i < len(feature.ChildrenBehind); i++ {
			drawFeature(&feature.ChildrenBehind[i], tt)
		}
	}
	if feature.Texture != nil {
		op.GeoM.Scale(3, 3)
		characterBuffer.DrawImage(feature.Texture, &op)
	}
	if feature.ChildrenBefore != nil {
		for i := 0; i < len(feature.ChildrenBefore); i++ {
			drawFeature(&feature.ChildrenBefore[i], tt)
		}
	}
}

func DrawCharacter(ch *world.Character, world *world.Model, screen *ebiten.Image) {

	characterBuffer.Fill(color.RGBA{10, 10, 10, 25})

	drawFeature(&ch.Appearance.Body, ebiten.GeoM{})

	xf, yf := world.Camera.ToScreen(ch.Pos).ValuesFloat()

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(xf, yf)

	screen.DrawImage(characterBuffer, &op)

}
