package main

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/suapapa/go_devices/epd7in5v2"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

var (
	dev *epd7in5v2.Dev
)

func initHW() error {
	if flagDryrun {
		return nil
	}

	if _, err := host.Init(); err != nil {
		return err
	}

	s, err := spireg.Open("")
	if err != nil {
		return err
	}

	dev, err = epd7in5v2.NewSPIHat(s)
	if err != nil {
		return err
	}
	return nil
}

func updatePanel(img image.Image) {
	img = imaging.Rotate(img, 0, color.White)

	if err := dev.Draw(img.Bounds(), img, image.Point{}); err != nil {
		panic(err)
	}
}
