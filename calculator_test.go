package main

import (
	"fmt"
	"image"
	"math/big"
	"testing"
)

func describe(desc string, expected, actual interface{}) string {
	return fmt.Sprintf("%s\nExpected: %v\nbut got: %v", desc, expected, actual)
}

var calcRatioTests = []struct {
	width, height int
	name          string
}{
	{100, 100, "square"},
	{4032, 3024, "landscape < 1.5"},
	{6000, 4000, "landscape == 1.5"},
	{3867, 2188, "landscape > 1.5"},
	{1711, 2429, "portrait < 1.5"},
	{4000, 6000, "portrait == 1.5"},
	{1711, 2829, "portrait > 1.5"},
}

func TestMatCalculator(t *testing.T) {

	expected := "1.5"

	for _, tt := range calcRatioTests {

		mw, mh := calcMat(tt.width, tt.height)
		var actual string
		if mh > mw {
			actual = big.NewRat(int64(mh), int64(mw)).FloatString(1)
		} else {
			actual = big.NewRat(int64(mw), int64(mh)).FloatString(1)
		}

		if actual != expected {
			t.Error(describe("wrong ratio; "+tt.name, expected, actual))
		}

	}

}

var offsetTests = []struct {
	w, h       int // input
	xoff, yoff int // expected
	mat        image.Rectangle
	name       string
}{
	{60, 10, 0, -15, landscape, "y"},
	{20, 60, -10, 0, portrait, "x"},
	{15, 60, -12, 0, portrait, "x"},
}

var landscape = image.Rect(0, 0, 60, 40)
var portrait = image.Rect(0, 0, 40, 60)

func Test_matOffset(t *testing.T) {

	for _, tt := range offsetTests {

		pic := image.Rect(0, 0, tt.w, tt.h)

		expected := image.Pt(tt.xoff, tt.yoff)

		actual := calcMatOffset(tt.mat, pic)

		if actual != expected {
			t.Error(describe("wrong "+tt.name+" offset", expected, actual))
		}

	}

}
