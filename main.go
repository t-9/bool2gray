/*
bool2gray converts a two-dimensional bool slice to a gray scale image.
*/
package bool2gray

import (
	"image"
	"image/color"
)

// Do converts a bool 2-dim slice to a gray scale image.
func Do(src [][]bool) (img *image.Gray) {
	w := 0
	for _, l := range src {
		if w < len(l) {
			w = len(l)
		}
	}

	img = image.NewGray(image.Rect(0, 0, w, len(src)))
	for y, l := range src {
		for x, f := range l {
			if f {
				img.Set(x, y, color.White)
			}
		}
	}
	return
}

// Enlarge converts a two-dimensional bool slice to a gray scale image
// enlarged to a power of two.
func Enlarge(src [][]bool, e int) (img *image.Gray) {
	b := make([][]bool, len(src))
	for y, l := range src {
		b[y] = make([]bool, len(l))
	}

	for y, l := range src {
		for x, f := range l {
			b[y][x] = f
		}
	}

	for i := 0; i < e; i++ {
		b = enlargeX2(b)
	}
	img = Do(b)
	return
}

func enlargeX2(src [][]bool) (b [][]bool) {
	b = make([][]bool, len(src)*2)
	for y := range src {
		ty := y * 2
		b[ty] = make([]bool, len(src[y])*2)
		b[ty+1] = make([]bool, len(src[y])*2)
	}

	for y, l := range src {
		for x, f := range l {
			tx := x * 2
			ty := y * 2
			b[ty][tx] = f
			b[ty][tx+1] = f
			b[ty+1][tx] = f
			b[ty+1][tx+1] = f
		}
	}
	return
}
