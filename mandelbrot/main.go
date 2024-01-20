package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width   = 800
	height  = 800
	maxIter = 1000
)

func mandelbrot(c complex128) color.Color {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return color.RGBA{uint8(i % 256), uint8((i * 2) % 256), uint8((i * 5) % 256), 255}
		}
	}
	return color.Black
}

func render() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	minReal, maxReal := -2.0, 2.0
	minImag, maxImag := -2.0, 2.0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			realPart := minReal + (maxReal-minReal)*float64(x)/float64(width-1)
			imagPart := minImag + (maxImag-minImag)*float64(y)/float64(height-1)
			c := complex(realPart, imagPart)
			tone := mandelbrot(c)
			img.Set(x, y, tone)
		}
	}

	return img
}

func saveImg(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func main() {
	img := render()
	saveImg(img, "mandelbrot.png")
}
