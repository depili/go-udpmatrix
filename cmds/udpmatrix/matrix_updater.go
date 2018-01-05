package main

import (
	"github.com/mcuadros/go-rpi-rgb-led-matrix"
	"image"
	"image/color"
	"image/draw"
	"time"
)

func initMatrix() chan byte {
	config := &rgbmatrix.DefaultConfig
	config.Rows = options.Rows
	config.Parallel = options.Parallel
	config.ChainLength = options.Chain
	config.Brightness = options.Brightness

	c := make(chan byte, 50)

	go runMatrix(config, c)
	return c
}

func runMatrix(config *rgbmatrix.HardwareConfig, c chan byte) {
	m, err := rgbmatrix.NewRGBLedMatrix(config)
	fatal(err)

	canvas := rgbmatrix.NewCanvas(m)
	defer canvas.Close()

	bounds := canvas.Bounds()
	img := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	x := 0
	y := 0

	var red, green, blue, color_channel uint8
	ticker := time.NewTicker(time.Millisecond * 100)

	fmt.Printf("Matrix updater running\n")

	for {
		select {
		case <-ticker.C:
			draw.Draw(canvas, canvas.Bounds(), img, image.ZP, draw.Src)
			canvas.Render()
		case b := <-c:
			switch color_channel {
			case 0:
				red = uint8(b)
			case 1:
				green = uint8(b)
			case 2:
				blue = uint8(b)
				color := color.RGBA{red, green, blue, 255}
				img.Set(x, y, color)
				color_channel = 0
				// Select the next pixel in loop
				x++
				if x >= bounds.Dx() {
					x = 0
					y++
					if y >= bounds.Dy() {
						y = 0
					}
				}
			}
		}
	}
	fmt.Printf("Matrix updater done.\n")
}
