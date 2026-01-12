package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func VisualizeNumberArray(screen *ebiten.Image, numbers []int, current int, next int) {
	offset := 1
	activeIndexStrokeWidth := 2
	nextIndexStrokeWidth := 4
	heightMultiplier := 30

	for i, number := range numbers {
		height := float32(number * heightMultiplier)

		x := float32((i + offset) * (rectangleWidth + 2))
		y := float32(rectangleHeight - int(height))

		vector.FillRect(screen, x, y, float32(rectangleWidth), height, color.RGBA{0xff, 0, 0, 0xff}, true)

		if current == i {
			vector.StrokeRect(screen, x, y, float32(rectangleWidth), height, float32(activeIndexStrokeWidth), color.White, true)

			if i+1 < len(numbers) {
				nextHeight := float32(numbers[i+1] * heightMultiplier)

				z := float32((i + 1 + offset) * (rectangleWidth + 2))
				k := float32(rectangleHeight - int(nextHeight))

				vector.StrokeRect(screen, z, k, float32(rectangleWidth), nextHeight, float32(nextIndexStrokeWidth), color.RGBA{0xff, 0xff, 0, 0xff}, true)
			}
		}

		numberString := fmt.Sprintf("%d", number)
		textX := x + float32(rectangleWidth)/2
		textY := y + height/2 - float32(fontSize)

		textOptions := &text.DrawOptions{}
		textOptions.GeoM.Translate(float64(int(textX)-4), float64(int(textY)+4))
		textOptions.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, numberString, &text.GoTextFace{
			Source: pressStartSource,
			Size:   fontSize,
		}, textOptions)
	}
}
