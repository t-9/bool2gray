package bool2gray

import (
	"image"
	"image/png"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	in := [][]bool{
		{true, false},
		{false, true},
	}

	expected := &image.Gray{
		Pix:    []uint8{0xff, 0x0, 0x0, 0xff},
		Stride: 2,
		Rect: image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: 2,
				Y: 2,
			},
		},
	}

	actueal := Do(in)

	if !reflect.DeepEqual(&actueal, &expected) {
		t.Fatalf("failed test\nexpected: %#v\nactual: %#v", expected, actueal)
	}
}

func TestEnlarge(t *testing.T) {
	in := [][]bool{
		{true, false},
		{false, true},
	}
	e := 1

	expected := &image.Gray{
		Pix: []uint8{
			0xff, 0xff, 0x0, 0x0,
			0xff, 0xff, 0x0, 0x0,
			0x0, 0x0, 0xff, 0xff,
			0x0, 0x0, 0xff, 0xff,
		},
		Stride: 4,
		Rect: image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: 4,
				Y: 4,
			},
		},
	}

	actueal := Enlarge(in, e)

	if !reflect.DeepEqual(&actueal, &expected) {
		t.Fatalf("failed test\nexpected: %#v\nactual: %#v", expected, actueal)
	}
}

func TestEnlargeX2(t *testing.T) {
	in := [][]bool{
		{true, false},
		{false, true},
	}

	expected := [][]bool{
		{true, true, false, false},
		{true, true, false, false},
		{false, false, true, true},
		{false, false, true, true},
	}

	actueal := enlargeX2(in)

	if !reflect.DeepEqual(actueal, expected) {
		t.Fatalf("failed test\nexpected: %#v\nactual: %#v", expected, actueal)
	}
}

func ExampleDo() {
	in := [][]bool{
		{true, false},
		{false, true},
	}

	img := Do(in)

	file, err := os.Create("sample.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}

func ExampleEnlarge() {
	in := [][]bool{
		{true, false},
		{false, true},
	}
	e := 5

	img := Enlarge(in, e)

	file, err := os.Create("sample.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}
