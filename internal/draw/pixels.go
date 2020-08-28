package draw

import "image/color"

func SetPixel(x int, y int, w int, c color.RGBA, dst []byte) {
	i := y*w*4 + x*4
	dst[i] = c.R
	dst[i+1] = c.G
	dst[i+2] = c.B
	dst[i+3] = c.A
}
