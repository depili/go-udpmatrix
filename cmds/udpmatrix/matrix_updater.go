package main

import (
	"github.com/mcuadros/go-rpi-rgb-led-matrix"
	"image/color"
)

func initMatrix() chan byte {
	config := &rgbmatrix.DefaultConfig
	config.Rows = *options.Rows
	config.Parallel = *options.Parallel
	config.ChainLength = *options.Chain
	config.Brightness = *options.Brightness

	c = make(chan byte)

	go runMatrix(config, c)
	return c
}

func runMatrix(config *rgbmatrix.HardwareConfig, c chan byte) {
	m, err := rgbmatrix.NewRGBLedMatrix(config)
	fatal(err)

	canvas := rgbmatrix.NewCanvas(m)
	defer canvas.Close()

	bounds := canvas.Bounds()
	x := 0
	y := 0

	var red, green, blue, color_channel uint16

	for b := range c {
		switch color_channel {
		case 0:
			red = uint16(b)
		case 1:
			green = uint16(b)
		case 2:
			blue = uint16(b)
			color := color.RGBA{red, green, blue, 255}
			canvas.Set(x, y, color)
			canvas.Render()

			// Select the next pixel in loop
			x++
			if x > bounds.Dx() {
				x = 0
				y++
				if y > bounds.Dy() {
					y = 0
				}
			}
		}
		color_channel = (color_channel + 1) % 3
	}
}
