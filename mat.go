package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strings"
)

// 1. calc size with mat
// 2. calc matt offset
// 	- (matt - orig size) / 2
// 3. create image the size with mat
// 4. copy photo image to mat, starting at matt offset

type Foo struct {
	src  image.Image
	dst  image.Image
	xoff int
	yoff int
}

func mat(f Foo) {
	// image.NewNRGBA(j
	// draw.Draw(f.dst, f.src.Bounds(), f.src, image.Pt(xoff, yoff), draw.Src)
}

func main() {

	pic, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	im, _, err := image.Decode(pic)
	if err != nil {
		panic(err)
	}

	mw, mh := calcMat(im.Bounds().Dx(), im.Bounds().Dy())

	mat := image.Rect(0, 0, mw, mh)
	dst := image.NewNRGBA(mat)
	draw.Draw(dst, mat, &image.Uniform{color.White}, image.ZP, draw.Src)

	offset := calcMatOffset(mat, im.Bounds())
	draw.Draw(dst, mat, im, offset, draw.Src)

	matted, err := os.Create(name(pic.Name()))
	if err != nil {
		panic(err)
	}

	opts := &jpeg.Options{100}
	err = jpeg.Encode(matted, dst, opts)
	if err != nil {
		panic(err)
	}
}

func name(filename string) string {
	lastPeriod := strings.LastIndex(filename, ".")
	return filename[:lastPeriod] + "_mat_4x6" + filename[lastPeriod:]
}
