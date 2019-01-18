package main

import (
	"image"
	"image/color"
)

// Bool2Gray converts a bool 2-dim slice to a gray image.
func Bool2Gray(src [][]bool) (img *image.Gray) {
	h := len(src)
	if h == 0 {
		return
	}
	w := len(src[0])
	img = image.NewGray(image.Rect(0, 0, w, h))

	black := color.Gray{0xff}
	for y, yf := range src {
		for x, xf := range yf {
			if xf {
				img.Set(x, y, black)
			}
		}
	}
	return
}
