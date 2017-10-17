package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var profile = flag.Bool("p", false, "write cpu profile `file`")

func main() {

	flag.Parse()
	if *profile {

		log.Println("pprof enabled")

		f, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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

	bufmat := bufio.NewWriterSize(matted, 8192)

	opts := &jpeg.Options{100}
	err = jpeg.Encode(bufmat, dst, opts)
	if err != nil {
		panic(err)
	}
}

func name(filename string) string {
	lastPeriod := strings.LastIndex(filename, ".")
	return filename[:lastPeriod] + "_mat_4x6" + filename[lastPeriod:]
}
