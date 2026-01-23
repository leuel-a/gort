package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	offset                 = 2
	activeIndexStrokeWidth = 2
	nextIndexStrokeWidth   = 4
	heightMultiplier       = 10
	spacerWidth            = 2
)

func VisualizeNumberArray(screen *ebiten.Image, numbers []int, current int, next int) {
	n := len(numbers)

	totalWidth := n*rectangleWidth + (n-1)*spacerWidth
	startX := (baseWidth - totalWidth) / 2

	for i, number := range numbers {
		height := float32(number * heightMultiplier)

		x := float32(startX + i*(rectangleWidth+spacerWidth))
		y := float32(rectangleHeight - int(height))

		vector.FillRect(screen, x, y, float32(rectangleWidth), height, color.RGBA{0xff, 0, 0, 0xff}, true)

		if current == i {
			vector.StrokeRect(screen, x, y, float32(rectangleWidth), height,
				float32(activeIndexStrokeWidth), color.White, true)

			if i+1 < n {
				nextHeight := float32(numbers[i+1] * heightMultiplier)
				nx := float32(startX + (i+1)*(rectangleWidth+spacerWidth))
				ny := float32(rectangleHeight - int(nextHeight))

				vector.StrokeRect(screen, nx, ny, float32(rectangleWidth), nextHeight,
					float32(nextIndexStrokeWidth), color.RGBA{0xff, 0xff, 0, 0xff}, true)
			}
		}

		textX := x + float32(rectangleWidth)/2
		textY := y + height/2 - float32(fontSize)

		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(textX-4), float64(textY+4))
		opts.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, fmt.Sprintf("%d", number), &text.GoTextFace{
			Source: poppinsSource,
			Size:   fontSize,
		}, opts)
	}
}

func VisualizeNumberArraySorted(screen *ebiten.Image, numbers []int) {
	n := len(numbers)

	totalWidth := n*rectangleWidth + (n-1)*spacerWidth
	startX := (baseWidth - totalWidth) / 2

	for i, number := range numbers {
		height := float32(number * heightMultiplier)

		x := float32(startX + i*(rectangleWidth+spacerWidth))
		y := float32(rectangleHeight - int(height))

		vector.FillRect(screen, x, y, float32(rectangleWidth), height,
			color.RGBA{0xff, 0, 0, 0xff}, true)

		vector.StrokeRect(screen, x, y, float32(rectangleWidth), height,
			float32(activeIndexStrokeWidth), color.RGBA{0, 0xff, 0, 0xff}, true)

		textX := x + float32(rectangleWidth)/2
		textY := y + height/2 - float32(fontSize)

		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(textX-4), float64(textY+4))
		opts.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, fmt.Sprintf("%d", number), &text.GoTextFace{
			Source: poppinsSource,
			Size:   fontSize,
		}, opts)
	}
}
