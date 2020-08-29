package draw

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
)

func PlayerSprite(ch *world.Character, w *world.Model, screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	sprite := textures.SpriteAlycia.Image()
	if ch.Direction < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(16, 0)
	}

	op.GeoM.Scale(3, 3)

	i := 0
	j := (ch.Step / 30) % 6

	if ch.Vel.X != 0 {
		i = 1
		j = (ch.Step / 7) % 6
	}

	if ch.Pos.Y < 0 {
		f := (ch.Step / 6) % 2
		i = 2
		if ch.Vel.Y < 0 {
			j = 0 + f
		} else if ch.Vel.Y >= 3 {
			j = 4 + f
		} else if ch.Vel.Y >= 0 {
			j = 2 + f
		}
	}

	xf, yf := w.Camera.ToScreen(ch.Pos).ValuesFloat()

	op.GeoM.Translate(xf, yf)

	// create an ebiten image as textures are imported as rgba go images
	img2, _ := ebiten.NewImageFromImage(sprite.SubImage(image.Rect(16*j, i*28, 16*(j+1), (i+1)*28)), ebiten.FilterDefault)
	_ = screen.DrawImage(img2, op)
}
