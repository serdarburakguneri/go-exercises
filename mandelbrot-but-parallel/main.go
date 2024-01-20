package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

const (
	width        = 1600
	height       = 1600
	maxIter      = 1000
	routineCount = 8
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

func renderRow(waitGroup *sync.WaitGroup, img *image.RGBA, yMin, yMax int) {
	defer waitGroup.Done()
	minReal, maxReal := -2.0, 2.0
	minImag, maxImag := -2.0, 2.0

	for y := yMin; y < yMax; y++ {
		for x := 0; x < width; x++ {
			realPart := minReal + (maxReal-minReal)*float64(x)/float64(width-1)
			imagPart := minImag + (maxImag-minImag)*float64(y)/float64(height-1)
			c := complex(realPart, imagPart)
			tone := mandelbrot(c)
			img.Set(x, y, tone)
		}
	}
}

func render() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var waitGroup sync.WaitGroup
	waitGroup.Add(routineCount)

	// this time, instead of rendering the whole image, we render rows in parallel

	rowPerRoutine := height / routineCount

	for j := 0; j < routineCount; j++ {

		yMin := j * rowPerRoutine
		yMax := (j + 1) * rowPerRoutine

		go renderRow(&waitGroup, img, yMin, yMax)
	}

	waitGroup.Wait()

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
