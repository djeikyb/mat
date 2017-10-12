package main

import (
	"image"
	"math/big"
	"strconv"
)

func calcMatOffset(mat, pic image.Rectangle) image.Point {
	if mat.Dx() > pic.Dx() && mat.Dy() > pic.Dy() {
		panic("mat is both wider and taller than photo")
	}
	if mat.Dx() < pic.Dx() && mat.Dy() < pic.Dy() {
		panic("mat is both narrower and shorter than photo")
	}
	if mat.Dx() != pic.Dx() {
		return image.Pt(-1*(mat.Dx()-pic.Dx())/2, 0)
	}
	return image.Pt(0, -1*(mat.Dy()-pic.Dy())/2)
}

func calcMat(w, h int) (int, int) {

	if w < h {
		return calcPortrait(w, h)
	}

	// squares will have mat on left and right
	return calcLandscape(w, h)
}

func calcPortrait(w, h int) (int, int) {

	ratio := big.NewRat(int64(h), int64(w))
	sixfour := big.NewRat(6, 4)

	c := ratio.Cmp(sixfour)

	if c == 0 {
		return w, h
	}

	// need mat above and below
	if c < 0 {
		mh := big.NewRat(int64(6*w), 4)
		return w, atoi(mh.FloatString(0))
	}

	// need mat left and right
	mw := big.NewRat(int64(4*h), 6)
	return atoi(mw.FloatString(0)), h
}

func calcLandscape(w, h int) (int, int) {

	ratio := big.NewRat(int64(w), int64(h))
	sixfour := big.NewRat(6, 4)

	c := ratio.Cmp(sixfour)

	if c == 0 {
		return w, h
	}

	// need mat above and below
	if c > 0 {
		mh := big.NewRat(int64(4*w), 6)
		return w, atoi(mh.FloatString(0))
	}

	// need mat left and right
	mw := big.NewRat(int64(6*h), 4)
	return atoi(mw.FloatString(0)), h
}

func atoi(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
