package draw

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/ojrac/opensimplex-go"
	"image/color"
	"rings-of-ara/internal/world"
)

var treePixel *ebiten.Image

var n = opensimplex.NewNormalized(0)

func init() {
	treePixel, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
	treePixel.Fill(color.RGBA{
		R: 80,
		G: 80,
		B: 60,
		A: 255,
	})
}

func BackgroundLayer(w *world.Model, dst *ebiten.Image) {

	width := int(w.Camera.Size.W)
	/*
		for i := 0; i < width; i++ {
			r := n.Eval2((float64(w.Camera.Subject.Pos.X)/50+float64(i)/30), 0)
			if r < 0.75 {
				continue
			}

			op := ebiten.DrawImageOptions{}
			op.GeoM.Scale(1, float64(w.Camera.Size.H))
			op.GeoM.Translate(float64(i), 0)
			op.ColorM.Scale(2, 2, 2, 1)
			dst.DrawImage(treePixel, &op)
		}


		for i := 0; i < width; i++ {
			r := n.Eval2(float64(w.Camera.Subject.Pos.X)/30+float64(i)/20, 0)
			if r < 0.75 {
				continue
			}

			op := ebiten.DrawImageOptions{}
			op.GeoM.Scale(10, float64(w.Camera.Size.H))
			op.GeoM.Translate(float64(i), 0)
			op.ColorM.Scale(1.6, 1.6, 1.6, 1)
			dst.DrawImage(treePixel, &op)
		}*/

	for i := 0; i < width; i++ {
		r := n.Eval2((float64(w.Camera.Subject.Pos.X)*0.25+float64(i))/25, 0)
		if r < 0.75 {
			continue
		}

		op := ebiten.DrawImageOptions{}
		op.ColorM.Scale(2.2, 2.2, 2.2, 0.1)
		op.GeoM.Scale(15, float64(w.Camera.Size.H))
		op.GeoM.Translate(float64(i), 0)
		dst.DrawImage(treePixel, &op)
	}

	for i := 0; i < width; i++ {
		r := n.Eval2((float64(w.Camera.Subject.Pos.X)*0.5+float64(i))/50, 0)
		if r < 0.75 {
			continue
		}

		op := ebiten.DrawImageOptions{}
		op.ColorM.Scale(1.7, 1.7, 1.7, 0.25)
		op.GeoM.Scale(15, float64(w.Camera.Size.H))
		op.GeoM.Translate(float64(i), 0)
		dst.DrawImage(treePixel, &op)
	}

	for i := 0; i < width; i++ {
		r := n.Eval2((float64(w.Camera.Subject.Pos.X)*0.75+float64(i))/150, 0)
		if r < 0.75 {
			continue
		}

		op := ebiten.DrawImageOptions{}
		op.ColorM.Scale(1, 1, 1, 0.5)
		op.GeoM.Scale(15, float64(w.Camera.Size.H))
		op.GeoM.Translate(float64(i), 0)
		dst.DrawImage(treePixel, &op)
	}

}
