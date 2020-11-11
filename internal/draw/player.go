package draw

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	"math"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
)

var characterArmBuffer *ebiten.Image

func init() {
	characterArmBuffer, _ = ebiten.NewImage(28*4, 28*4, ebiten.FilterDefault)
}

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
		if ch.Vel.X <= 4 {
			j = (ch.Step / 6) % 6
		} else {
			j = (ch.Step / 5) % 6
		}
	}

	if ch.Vel.Y != 0 {
		f := (ch.Step / 6) % 2
		i = 2
		if ch.Vel.Y > 0 {
			j = 0 + f
		} else if ch.Vel.Y <= -3 {
			j = 4 + f
		} else {
			j = 2 + f
		}
	}

	xf, yf := w.Camera.ToScreen(ch.Pos).ValuesFloat()

	op.GeoM.Translate(xf, yf)

	if ch.Direction < 0 {
		i = 3
		j = 2
	} else {
		i = 3
		j = 0
	}

	// create an ebiten image as textures are imported as rgba go images
	img2, _ := sprite.SubImage(image.Rect(16*j, i*28, 16*(j+1), (i+1)*28)).(*ebiten.Image)

	// draw aim
	aimLine(ch, w.Camera, screen)

	if ch.Direction < 0 {
		j = 3
		i = 3
	} else {
		j = 1
		i = 3
	}
	_ = characterArmBuffer.Clear()

	ox, oy := w.Camera.ToScreen(ch.Pos).ValuesFloat()
	tx, ty := w.Camera.ToScreen(ch.Target).ValuesFloat()

	op4 := ebiten.DrawImageOptions{}
	op4.GeoM.Translate(-14, -8)
	op4.GeoM.Rotate(math.Atan2(ty-oy, tx-ox) - math.Pi/2)
	op4.GeoM.Translate(28*2, 28*2)
	wea := textures.TileSetWeapons.SubImage(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{28, 28},
	}).(*ebiten.Image)
	_ = characterArmBuffer.DrawImage(wea, &op4)

	op2 := ebiten.DrawImageOptions{}

	op2.GeoM.Translate(-5, -15)
	op2.GeoM.Rotate(math.Atan2(ty-oy, tx-ox) - math.Pi/2)
	op2.GeoM.Translate(28*2, 28*2)

	img3, _ := sprite.SubImage(image.Rect(16*j, i*28, 16*(j+1), (i+1)*28)).(*ebiten.Image)
	_ = characterArmBuffer.DrawImage(img3, &op2)

	op3 := ebiten.DrawImageOptions{}
	op3.GeoM.Translate(-28*2, -28*2)
	op3.GeoM.Translate(5, 15)
	op3.GeoM.Scale(3, 3)
	op3.GeoM.Translate(xf, yf)


	if ch.Direction < 0 {
		_ = screen.DrawImage(characterArmBuffer, &op3)
		_ = screen.DrawImage(img2, op)
	} else {
		_ = screen.DrawImage(img2, op)
		_ = screen.DrawImage(characterArmBuffer, &op3)
	}

}

func aimLine(ch *world.Character, cam *world.Camera, dst *ebiten.Image) {

	ox, oy := cam.ToScreen(ch.Pos).ValuesFloat()
	tx, ty := cam.ToScreen(ch.Target).ValuesFloat()

	ebitenutil.DrawLine(dst, ox, oy, tx, ty, color.RGBA{0, 0, 0, 0})
}
