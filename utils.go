package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func drawSortingAlgorithmName(screen *ebiten.Image, game *Game, clr color.Color) {
	x := 0.0
	y := 100.0

	nameFontSize := 24.0
	fontFace := &text.GoTextFace{
		Source: pressStartSource,
		Size:   nameFontSize,
	}

	textW, _ := text.Measure(game.sortingAlgorithmName, fontFace, 0)
	textX := x + ((float64(screenWidth) - textW) / 2)

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(textX, float64(y))
	textOptions.ColorScale.ScaleWithColor(clr)

	text.Draw(screen, game.sortingAlgorithmName, fontFace, textOptions)
}

var prevMousePressedState bool

func mouseJustPressed() bool {
	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	justPressed := pressed && !prevMousePressedState

	prevMousePressedState = pressed
	return justPressed
}

func pointInRect(px, py, x, y, w, h float32) bool {
	return px >= x && px <= x+w && py >= y && py <= y+h
}
