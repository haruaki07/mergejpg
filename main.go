package main

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	args = args[1:]
	if len(args) < 3 {
		println("Usage:\n    mergerapor [image1] [image2] [output]")
		os.Exit(0)
	}

	img1File, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}

	img2File, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}

	img1, _, err := image.Decode(img1File)
	if err != nil {
		panic(err)
	}

	img2, _, err := image.Decode(img2File)
	if err != nil {
		panic(err)
	}

	resize := 30 // 30%
	if len(args) == 4 {
		resize, err = strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
	}

	result := Merge(img1, img2)
	result = Resize(result, resize)

	out, err := os.Create(args[2])
	if err != nil {
		panic(err)
	}

	var opt jpeg.Options
	opt.Quality = 100

	jpeg.Encode(out, result, &opt)
	println("success!")
}
